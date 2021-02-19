package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthChecker is a function to check whether the app is working properly or not
func HealthChecker(e echo.Context) error {
	response := struct {
		Success bool `json:"success"`
	}{
		true,
	}
	return e.JSON(http.StatusOK, response)
}
