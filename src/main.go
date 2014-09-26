package main

import (
	"io"
	"mantis"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "Welcome home")
}

func login(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "please login")
}

func main() {
	m := mantis.Classic()

	m.AddRouter("/", home)

	m.AddRouter("/login", login)

	m.Run()
}
