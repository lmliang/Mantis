package mantis

import (
	"net/http"
	"reflect"
	"strings"
)

type Route interface {
	Match(pattern string, method string) bool
	Handle(rw http.ResponseWriter, r *http.Request)
}

type route struct {
	pattern string
	method  string
	handler Handler
}

func (rt *route) Match(pattern string, method string) bool {
	if rt.pattern == pattern && (len(rt.method) <= 0 || rt.method == strings.ToUpper(method)) {
		return true
	} else {
		return false
	}
}

func (rt *route) Handle(rw http.ResponseWriter, r *http.Request) {
	t := typeOfPtr(rt.handler)
	v := reflect.ValueOf(rt.handler)

	if t.Kind() == reflect.Func {
		in := make([]reflect.Value, t.NumIn())
		in[0] = reflect.ValueOf(rw)
		in[1] = reflect.ValueOf(r)
		v.Call(in)
	} else if t.Kind() == reflect.Struct {
		if ctrl, ok := rt.handler.(Controller); ok {
			dispatchController(rw, r, ctrl)
		}
	}
}

func dispatchController(rw http.ResponseWriter, r *http.Request, c Controller) {
	ctx := newContext(rw, r)

	c.Before()

	ctrl := valueOfPtr(c)
	if fd := ctrl.FieldByName("Ctx"); fd.IsValid() && fd.CanSet() {
		fd.Set(reflect.ValueOf(ctx))

		switch strings.ToUpper(r.Method) {
		case "GET":
			c.Get()
		case "POST":
			c.Post()
		case "PUT":
			c.Put()
		case "PATCH":
			c.Patch()
		case "DELETE":
			c.Delete()
		case "HEAD":
			c.Head()
		case "OPTIONS":
			c.Options()
		}

		c.Render()

		c.Finish()

	} else {
		panic("Dispatch Controller Failed.")
	}
}

func newRoute(method string, pattern string, h Handler) Route {
	t := typeOfPtr(h)

	if t.Kind() == reflect.Struct {
		ctrl := valueOfPtr(h)
		if fd := ctrl.FieldByName("Data"); fd.IsValid() && fd.CanSet() {
			m := make(map[interface{}]interface{})
			fd.Set(reflect.ValueOf(m))
		}
	}

	return &route{method: method, pattern: pattern, handler: h}
}
