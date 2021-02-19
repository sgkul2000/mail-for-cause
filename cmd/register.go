package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
}
