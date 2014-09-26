package mantis

import (
	"fmt"
	"net/http"
)

type Controller interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
	Head()
	Options()
	Redirect(string, int)
}

type DefaultController struct {
	Ctx  *Context
	Tmpl string
	Data map[interface{}]interface{}
}

func (c *DefaultController) Get() {
	fmt.Println("Enter Controller Get")
	http.Error(c.Ctx.Resp, "Method Not Allowed", 405)
}

func (c *DefaultController) Post() {
	http.Error(c.Ctx.Resp, "Method Not Allowed", 405)
}

func (c *DefaultController) Put() {
	http.Error(c.Ctx.Resp, "Method Not Allowed", 405)
}

func (c *DefaultController) Patch() {
	http.Error(c.Ctx.Resp, "Method Not Allowed", 405)
}

func (c *DefaultController) Delete() {
	http.Error(c.Ctx.Resp, "Method Not Allowed", 405)
}

func (c *DefaultController) Head() {
	http.Error(c.Ctx.Resp, "Method Not Allowed", 405)
}

func (c *DefaultController) Options() {
	http.Error(c.Ctx.Resp, "Method Not Allowed", 405)
}

func (c *DefaultController) Redirect(url string, code int) {
	c.Ctx.Redirect(url, code)
}
