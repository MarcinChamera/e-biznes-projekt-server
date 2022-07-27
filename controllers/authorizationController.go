package controllers

import (
	"backend/helpers"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

var googleConfig = &oauth2.Config{
	RedirectURL:  "https://namelessshop-server.azurewebsites.net/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

var githubConfig = &oauth2.Config{
	RedirectURL:  "https://namelessshop-server.azurewebsites.net/github/callback",
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	Scopes: []string{
		"user:email",
		"read:user",
	},
	Endpoint: github.Endpoint,
}

var facebookConfig = &oauth2.Config{
	RedirectURL:  "https://namelessshop-server.azurewebsites.net/facebook/callback",
	ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
	ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
	Scopes: []string{
		"email",
	},
	Endpoint: facebook.Endpoint,
}

func GoogleCallback(c echo.Context) error {
	token, err := googleConfig.Exchange(context.Background(), c.QueryParam("code"))

	if err != nil {
		fmt.Printf("googleConfig.Exchange error %s", err)
		return err
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		fmt.Printf("http.get google ouath error %s", err)
		return err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("google ioutil.ReadAll error %s", err)
		return err
	}

	user := struct {
		Email string
	}{}

	err = json.Unmarshal(contents, &user)

	if err != nil {
		fmt.Printf("google json.Unmarshal error %s", err)
		return err
	}

	user.Email = strings.ToLower(user.Email)
	if !CheckIfUserExists(user.Email, "google") {
		AddUser(user.Email, "google", *token)
	}

	userFromGet := GetUser(user.Email, "google")

	newToken := helpers.GenerateNewToken(40)
	userId := strconv.Itoa(int(userFromGet.ID))

	c.Redirect(http.StatusFound, "https://namelessshop.azurewebsites.net/login/auth/google/success/"+newToken+"&"+user.Email+"&"+userId)

	return c.JSON(http.StatusOK, echo.Map{
		"token": newToken,
		"user":  userFromGet,
	})
}

func GoogleLogin(c echo.Context) error {
	url := googleConfig.AuthCodeURL("state")
	return c.JSON(http.StatusOK, url)
}

func GithubCallback(c echo.Context) error {
	token, err := githubConfig.Exchange(context.Background(), c.QueryParam("code"))

	if err != nil {
		fmt.Printf("githubConfig.Exchange error %s", err)
		return err
	}

	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		fmt.Println("http.NewRequest GET error")
		return err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+token.AccessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http.DefaultClient.Do error")
		print(err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("github ioutil.ReadAll error %s", err)
		return err
	}

	user := struct {
		Email string
	}{}

	err = json.Unmarshal(contents, &user)

	if err != nil {
		fmt.Printf("github json.Unmarshal error %s", err)
		return err
	}

	user.Email = strings.ToLower(user.Email)
	if !CheckIfUserExists(user.Email, "github") {
		AddUser(user.Email, "github", *token)
	}

	userFromGet := GetUser(user.Email, "github")

	newToken := helpers.GenerateNewToken(40)

	c.Redirect(http.StatusFound, "https://namelessshop.azurewebsites.net/login/auth/github/success/"+newToken+"&"+user.Email+"&"+strconv.Itoa(int(userFromGet.ID)))

	return c.JSON(http.StatusOK, echo.Map{
		"token": newToken,
		"user":  userFromGet,
	})
}

func GithubLogin(c echo.Context) error {
	url := githubConfig.AuthCodeURL("state")
	return c.JSON(http.StatusOK, url)
}

func FacebookCallback(c echo.Context) error {
	token, err := facebookConfig.Exchange(context.Background(), c.QueryParam("code"))

	if err != nil {
		fmt.Printf("facebookConfig.Exchange error %s", err)
		return err
	}

	req, err := http.NewRequest("GET", "https://graph.facebook.com/me?fields=email&access_token="+token.AccessToken, nil)
	if err != nil {
		fmt.Println("http.NewRequest GET error")
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("facebook http.DefaultClient.Do error")
		print(err)
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("facebook ioutil.ReadAll error %s", err)
		return err
	}

	user := struct {
		Email string
	}{}

	err = json.Unmarshal(contents, &user)

	if err != nil {
		fmt.Printf("facebook json.Unmarshal error %s", err)
		return err
	}

	user.Email = strings.ToLower(user.Email)
	if !CheckIfUserExists(user.Email, "facebook") {
		AddUser(user.Email, "facebook", *token)
	}

	userFromGet := GetUser(user.Email, "facebook")

	newToken := helpers.GenerateNewToken(40)

	c.Redirect(http.StatusFound, "https://namelessshop.azurewebsites.net/login/auth/facebook/success/"+newToken+"&"+user.Email+"&"+strconv.Itoa(int(userFromGet.ID)))

	return c.JSON(http.StatusOK, echo.Map{
		"token": newToken,
		"user":  userFromGet,
	})
}

func FacebookLogin(c echo.Context) error {
	url := facebookConfig.AuthCodeURL("state")
	return c.JSON(http.StatusOK, url)
}
