package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// EchoErrorHandler is An error handler for echo frameword
func EchoErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	// c.Echo().DefaultHTTPErrorHandler(err, c)
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.JSON(code, map[string]interface{}{
		"success": false,
		"error":   err.Error(),
	})
}
