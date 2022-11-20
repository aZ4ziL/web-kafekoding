package main

import "github.com/gin-contrib/multitemplate"

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("login", "views/users/login.tmpl")
	r.AddFromFiles("register", "views/users/register.tmpl")

	return r
}
