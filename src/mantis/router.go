package mantis

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)
type Handler interface{}

func typeOfPtr(i interface{}) reflect.Type {
	t := reflect.TypeOf(i)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}

func validHandler(i interface{}) bool {
	t := typeOfPtr(i)
	fmt.Println("validHandler: ", t)

	if t.Kind() == reflect.Func {
	} else if t.Kind() == reflect.Interface {
		if _, ok := i.(Handler); ok {
			return true
		}
	}

	return false
}

type Router interface {
	AddRouter(pattern string, h Handler, method string)
	Handle(rw http.ResponseWriter, r *http.Request)
	NotFound()
}

type router struct {
	routes []Route
}

func (rt *router) AddRouter(pattern string, h Handler, method string) {
	if validHandler(h) {
		r := newRoute(strings.ToUpper(method), pattern, h)
		rt.routes = append(rt.routes, r)
	} else {
		fmt.Println("Invlaid handler for pattern [", pattern, "]")
	}
}

func (rt *router) Handle(rw http.ResponseWriter, r *http.Request) {
	for _, route := range rt.routes {
		if route.Match(r.URL.Path, r.Method) {
			route.Handle(rw, r)
			return
		}
	}

	rt.NotFound()
}

func (rt *router) NotFound() {

}

func newRouter() Router {
	return &router{}
}
