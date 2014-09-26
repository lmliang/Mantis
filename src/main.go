package main

import (
	"fmt"
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
	fmt.Println("LoginCtrl Get")
	io.WriteString(c.Ctx.Resp, "Please Login")
}

func main() {
	m := mantis.Classic()

	m.AddRouter("/", home)

	m.AddRouter("/login", &LoginController{})

	m.Run()
}
