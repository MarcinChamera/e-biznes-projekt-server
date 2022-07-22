package controllers

import (
	"backend/database"
	"backend/database/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func CheckIfUserExists(email string, service string) bool {
	var user models.User
	database.GetDatabase().Find(&user, "Email = ? AND o_auth_service = ?", email, service)
	if user.Email == "" {
		return false
	}
	return true
}

func AddUser(email string, service string, token oauth2.Token) {
	user := new(models.User)
	user.Email = email
	user.OAuthService = service
	user.Token = token.AccessToken
	database.GetDatabase().Create(user)
}

func GetUser(email string, service string) models.User {
	var user models.User
	database.GetDatabase().Find(&user, "email = ? AND o_auth_service = ?", email, service)
	return user
}

func GetUsers(c echo.Context) error {
	var users []models.User
	database.GetDatabase().Find(&users)
	return c.JSON(http.StatusOK, users)
}
