package main

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())
	m.Use(martini.Static("assets"))

	m.Get("/", func(ren render.Render) {
		ren.HTML(200, "index", nil)
	})

	m.Run()
}
