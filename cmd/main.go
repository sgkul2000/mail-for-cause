package main

import (
	"fmt"
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sgkul2000/mail-for-cause/pkg/db"
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

	// e.Validator = &types.CustomValidator{Validator: validator.New()}
	e.Static("/static", "views/static")
	Register(e)
	e.GET("/", func(c echo.Context) error {
		// return c.Render(http.StatusOK, "base", struct {
		// 	Page string
		// }{
		// 	Page: "index",
		// })

		causes, err := db.GetCauses()
		if err != nil {
			return err
		}

		tmpl, err := template.New("Main").Funcs(template.FuncMap{
			"trim": func(description string) string {
				if len(description) > 50 {
					return description[:50] + " ..."
				}
				return description
			},
		}).ParseGlob("views/*.html")
		if err != nil {
			return err
		}
		return tmpl.ExecuteTemplate(c.Response().Writer, "base", struct {
			Page string
			Data interface{}
		}{
			Page: "index",
			Data: causes,
		})
	})
	e.GET("/signup", func(c echo.Context) error {
		// return c.Render(http.StatusOK, "base", struct {
		// 	Page string
		// }{
		// 	Page: "index",
		// })
		tmpl, err := template.New("Main").Funcs(template.FuncMap{
			"trim": func(description string) string {
				if len(description) > 50 {
					return description[:50] + " ..."
				}
				return description
			},
		}).ParseGlob("views/*.html")
		if err != nil {
			return err
		}
		return tmpl.ExecuteTemplate(c.Response().Writer, "base", struct {
			Page string
		}{
			Page: "signup",
		})
	})
	e.GET("/login", func(c echo.Context) error {
		// return c.Render(http.StatusOK, "base", struct {
		// 	Page string
		// }{
		// 	Page: "index",
		// })
		tmpl, err := template.New("Main").Funcs(template.FuncMap{
			"trim": func(description string) string {
				if len(description) > 50 {
					return description[:50] + " ..."
				}
				return description
			},
		}).ParseGlob("views/*.html")
		if err != nil {
			return err
		}
		return tmpl.ExecuteTemplate(c.Response().Writer, "base", struct {
			Page string
		}{
			Page: "login",
		})
	})
	e.GET("/create", func(c echo.Context) error {
		// return c.Render(http.StatusOK, "base", struct {
		// 	Page string
		// }{
		// 	Page: "index",
		// })
		tmpl, err := template.New("Main").Funcs(template.FuncMap{
			"trim": func(description string) string {
				if len(description) > 50 {
					return description[:50] + " ..."
				}
				return description
			},
		}).ParseGlob("views/*.html")
		if err != nil {
			return err
		}
		return tmpl.ExecuteTemplate(c.Response().Writer, "base", struct {
			Page string
		}{
			Page: "create",
		})
	})
	e.GET("/cause/:name", func(c echo.Context) error {
		cause, err := db.GetCause(c.Param("name"))
		if err != nil {
			return err
		}
		// return c.Render(http.StatusOK, "base", struct {
		// 	Page string
		// }{
		// 	Page: "index",
		// })
		tmpl, err := template.New("Main").Funcs(template.FuncMap{
			"trim": func(description string) string {
				if len(description) > 50 {
					return description[:50] + " ..."
				}
				return description
			},
		}).ParseGlob("views/*.html")
		if err != nil {
			return err
		}
		fmt.Println(cause)
		return tmpl.ExecuteTemplate(c.Response().Writer, "base", struct {
			Page string
			Data interface{}
		}{
			Page: "send",
			Data: cause,
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
