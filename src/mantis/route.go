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
	t := reflect.TypeOf(rt.handler)
	v := reflect.ValueOf(rt.handler)

	if t.Kind() == reflect.Func {
		in := make([]reflect.Value, t.NumIn())
		in[0] = reflect.ValueOf(rw)
		in[1] = reflect.ValueOf(r)
		v.Call(in)
	} else if t.Kind() == reflect.Interface {
		if _, ok := rt.handler.(Handler); ok {

		}
	} else {

	}
}

func newRoute(method string, pattern string, h Handler) Route {
	return &route{method: method, pattern: pattern, handler: h}
}
