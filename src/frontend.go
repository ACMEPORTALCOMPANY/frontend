package main

import (
	"html/template"

	"github.com/ACMEPORTALCOMPANY/frontend/internal/handler"
	"github.com/ACMEPORTALCOMPANY/frontend/render"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/dist", "/dist")
	e.Static("/img", "/img")
	e.Renderer = &render.Renderer{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.GET("/", handler.App)

	e.Logger.Fatal(e.Start(":8080"))
}
