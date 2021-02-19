package main

import (
	"html/template"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sgkul2000/mail-for-cause/pkg/middlewares"
)

func main() {
	t := &Template{
		templates: template.Must(template.New("Main").Funcs(template.FuncMap{
			"trim": func(description string) string {
				if len(description) > 50 {
					return description[:50] + " ..."
				}
				return description
			},
		}).ParseGlob("views/*.html")),
	}

	e := echo.New()

	e.Renderer = t
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${latency_human}\n",
	}))
	e.HTTPErrorHandler = middlewares.EchoErrorHandler
	e.Use(middleware.Recover())

	e.Static("/static", "views/static")
	Register(e)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "1323"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
