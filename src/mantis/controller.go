package mantis

import (
	"fmt"
	"html/template"
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
	Before()
	Finish()
	Render()
	Redirect(string, int)
}

type DefaultController struct {
	Ctx  *Context
	Tmpl string
	Data map[interface{}]interface{}
}

func (c *DefaultController) Get() {
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

func (c *DefaultController) Before() {
	c.Tmpl = ""
	c.Data = make(map[interface{}]interface{})
}

func (c *DefaultController) Finish() {

}

func (c *DefaultController) Render() {
	if len(c.Tmpl) > 0 {
		t, err := template.ParseFiles(c.Tmpl)
		if err == nil {
			err = t.Execute(c.Ctx.Resp, c.Data)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}

func (c *DefaultController) Redirect(url string, code int) {
	c.Ctx.Redirect(url, code)
}
