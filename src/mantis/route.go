package mantis

import (
	"fmt"
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
	fmt.Println("call Handle")
	t := typeOfPtr(rt.handler)
	v := reflect.ValueOf(rt.handler)

	if t.Kind() == reflect.Func {
		in := make([]reflect.Value, t.NumIn())
		in[0] = reflect.ValueOf(rw)
		in[1] = reflect.ValueOf(r)
		v.Call(in)
	} else if t.Kind() == reflect.Struct {
		if ctrl, ok := rt.handler.(Controller); ok {
			fmt.Println("call dispatchController")
			dispatchController(rw, r, ctrl)
		}
	}
}

func dispatchController(rw http.ResponseWriter, r *http.Request, c Controller) {
	ctx := newContext(rw, r)

	fmt.Println("dispatchController set ctx")

	ctrl := valueOfPtr(c)
	if fd := ctrl.FieldByName("Ctx"); fd.IsValid() && fd.CanSet() {
		fmt.Println("dispatchController set ctx")
		fd.Set(reflect.ValueOf(ctx))
	}

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
}

func newRoute(method string, pattern string, h Handler) Route {
	return &route{method: method, pattern: pattern, handler: h}
}
