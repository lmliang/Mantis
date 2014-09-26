package mantis

import (
	"net/http"
)

type Context struct {
	Input  map[string]string
	Output map[string]string
	Req    *http.Request
	Resp   http.ResponseWriter
}

func (c *Context) Redirect(url string, code int) {
	c.Resp.Header().Set("Location", url)
	c.Resp.WriteHeader(code)
}

func (c *Context) parseForm() {
	c.Req.ParseForm()

	for k, _ := range c.Req.Form {
		c.Input[k] = c.Req.Form.Get(k)
	}
}

func (c *Context) Get(key string) string {
	if v, ok := c.Input[key]; ok {
		return v
	}

	return ""
}

func (c *Context) Set(key, val string) {
	c.Output[key] = val
}

func newContext(rw http.ResponseWriter, r *http.Request) *Context {
	in := make(map[string]string)
	out := make(map[string]string)
	ctx := &Context{Input: in, Output: out, Req: r, Resp: rw}

	ctx.parseForm()

	return ctx
}
