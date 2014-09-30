package main

import (
	"controller"
	"mantis"
)

func main() {
	m := mantis.Classic()

	m.AddRouter("/", controller.Home)

	m.AddRouter("/login", &controller.LoginController{})

	m.Run()
}
