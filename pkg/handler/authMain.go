package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sgkul2000/mail-for-cause/pkg/db"
	"github.com/sgkul2000/mail-for-cause/pkg/types"
)

// HandleRegister function is used to log a user in and generate a jwt token
func HandleRegister(c echo.Context) error {
	u := new(types.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	user, err := db.FindUser(u.Email)

	if err != nil {
		return err
	}
	if user.Email != "" {
		return echo.NewHTTPError(http.StatusBadRequest, "User already exists. Please login.")
	}
	id, err := db.CreateUser(*u)
	if err != nil {
		return err
	}
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"name":  u.Name,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 120).Unix(),
	}
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return err
	}
	response := types.Response{
		Success: true,
		Data: map[string]interface{}{
			"token":      tokenString,
			"InsertedID": id,
		},
	}
	return c.JSON(http.StatusOK, response)
}

// HandleLogin function is used to log a user in and generate a jwt token
func HandleLogin(c echo.Context) error {
	response := types.Response{
		Success: true,
	}
	u := new(types.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if err := validator.New().Struct(u); err != nil {
		return err
	}

	user, err := db.FindUser(u.Email)
	if err != nil {
		return err
	}
	result, err := u.ComparePasswords(user.Password)
	if err != nil {
		return err
	}
	if result == true {
		token := jwt.New(jwt.GetSigningMethod("HS256"))

		token.Claims = jwt.MapClaims{
			"username": user.Name,
			"email":    user.Email,
			"exp":      time.Now().Add(time.Hour * 120).Unix(),
		}
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			return err
		}
		response.Data = map[string]interface{}{
			"token": tokenString,
		}
		return c.JSON(http.StatusOK, response)
	}
	response.Success = false
	response.Error = "Wrong passwprd"
	return c.JSON(http.StatusUnauthorized, response)
}

// AuthChecker is used to chek whether the auth module is working or not
func AuthChecker(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Auth successful!",
	})
}
