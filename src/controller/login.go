package controller

import (
	"html/template"
	"io"
	"mantis"
	"net/http"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "Welcome home")
}

type LoginController struct {
	mantis.DefaultController
}

func (c *LoginController) Get() {
	t, err := template.ParseFiles("html/login.html")
	if err == nil {
		t.Execute(c.Ctx.Resp, nil)
		return
	} else {
		io.WriteString(c.Ctx.Resp, err.Error())
	}
}

func (c *LoginController) Post() {
	c.Tmpl = "html/home.html"
	c.Data["Name"] = c.Ctx.Get("username")
	c.Data["TimeNow"] = "2014-09-30 19:45:00"
}
