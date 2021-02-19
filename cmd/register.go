package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sgkul2000/mail-for-cause/pkg/db"
	"github.com/sgkul2000/mail-for-cause/pkg/handler"
)

// Register is a function to register alll the routes of the app
func Register(e *echo.Echo) {
	e.GET("/api/health-checker", handler.HealthChecker)
	e.POST("/api/login", handler.HandleLogin)
	e.POST("/api/register", handler.HandleRegister)

	// Routes with auth`
	e.GET("/api/auth-checker", handler.AuthChecker, middleware.JWT([]byte(os.Getenv("SECRET"))))

	e.GET("/api/cause", handler.GetCause, middleware.JWT([]byte(os.Getenv("SECRET"))))
	e.POST("/api/cause", handler.CreateCause, middleware.JWT([]byte(os.Getenv("SECRET"))))
	e.PUT("/api/cause", handler.CreateCause, middleware.JWT([]byte(os.Getenv("SECRET"))))

	// Views
	e.GET("/", func(c echo.Context) error {
		causes, err := db.GetCauses()
		if err != nil {
			return err
		}
		return c.Render(http.StatusOK, "base", struct {
			Page string
			Data interface{}
		}{
			Page: "index",
			Data: causes,
		})
	})
	e.GET("/signup", func(c echo.Context) error {
		return c.Render(http.StatusOK, "base", struct {
			Page string
		}{
			Page: "signup",
		})
	})
	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "base", struct {
			Page string
		}{
			Page: "login",
		})
	})
	e.GET("/create", func(c echo.Context) error {
		return c.Render(http.StatusOK, "base", struct {
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
		return c.Render(http.StatusOK, "base", struct {
			Page string
			Data interface{}
		}{
			Page: "send",
			Data: cause,
		})
	})
}
