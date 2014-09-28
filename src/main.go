package main

import (
	"html/template"
	"io"
	"mantis"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request) {
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
		io.WriteString(c.Ctx.Resp, "Please Login")
	}
}

func (c *LoginController) Post() {
	fmt.Println(c.Ctx.Input)
	c.Redirect("/", 301)
}

func main() {
	m := mantis.Classic()

	m.AddRouter("/", home)

	m.AddRouter("/login", &LoginController{})

	m.Run()
}
