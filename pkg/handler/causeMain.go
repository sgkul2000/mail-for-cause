package handler

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sgkul2000/mail-for-cause/pkg/db"
	"github.com/sgkul2000/mail-for-cause/pkg/types"
)

// CreateCause creates a new cause
func CreateCause(c echo.Context) error {
	tokenString := c.Request().Header.Values("Authorization")[0][7:]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return err
	}
	newCause := new(types.Cause)
	if err := c.Bind(newCause); err != nil {
		return err
	}
	if err = validator.New().Struct(newCause); err != nil {
		return err
	}
	err = db.CreateCause(claims["email"].(string), newCause)
	if err != nil {
		return err
	}
	response := types.Response{
		Success: true,
	}
	return c.JSON(http.StatusOK, response)
}

// EditCause creates a new cause
func EditCause(c echo.Context) error {
	tokenString := c.Request().Header.Values("Authorization")[0][7:]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return err
	}
	newCause := new(types.Cause)
	if err := c.Bind(newCause); err != nil {
		return err
	}
	if err = validator.New().Struct(newCause); err != nil {
		return err
	}
	err = db.EditCause(claims["email"].(string), newCause)
	if err != nil {
		return err
	}
	response := types.Response{
		Success: true,
	}
	return c.JSON(http.StatusOK, response)
}

// GetCause creates a new cause
func GetCause(c echo.Context) error {

	results, err := db.GetCauses()
	if err != nil {
		return err
	}
	response := types.Response{
		Success: true,
		Data:    results,
	}
	return c.JSON(http.StatusOK, response)
}

// // GetCauses creates a new cause
// func GetCauses(c echo.Context) error {
// 	tokenString := c.Request().Header.Values("Authorization")[0][7:]
// 	claims := jwt.MapClaims{}
// 	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("SECRET")), nil
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	newCause := new(types.Cause)
// 	if err := c.Bind(newCause); err != nil {
// 		return err
// 	}
// 	if err = validator.New().Struct(newCause); err != nil {
// 		return err
// 	}
// 	err = db.CreateCause(claims["email"].(string), newCause)
// 	if err != nil {
// 		return err
// 	}
// 	response := types.Response{
// 		Success: true,
// 	}
// 	return c.JSON(http.StatusOK, response)
// }
