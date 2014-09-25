package main

import (
	"io"
	"mantis"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "Welcome home")
}

func home2() {

}

func main() {
	m := mantis.Classic()

	m.AddRouter("/", home)

	//m.AddRouter("/home", home2)

	m.Run()
}
