package mantis

import (
	"net/http"
)

type Context struct {
	Req *http.Request
	RW  http.ResponseWriter
}

func newContext(rw http.ResponseWriter, r *http.Request) *Context {
	return &Context{Req: r, RW: rw}
}
