package handler

import (
	"net/http"

	"github.com/ACMEPORTALCOMPANY/frontend/internal/cookies"
	"github.com/ACMEPORTALCOMPANY/frontend/internal/panopticon"
	"github.com/labstack/echo/v4"
)

type AppTemplateData struct {
	IsLoggedIn bool
}

type ErrorData struct {
	Message string
}

func App(c echo.Context) error {
	data := AppTemplateData{}

	cookie, err := c.Cookie(cookies.TokenCookieName)
	if err != nil {
		data.IsLoggedIn = false
	} else {
		token, err := panopticon.LoginToken(cookie.Value)
		if err != nil {
			data.IsLoggedIn = false
		} else {
			newCookie := cookies.New(cookies.TokenCookieName, token)
			c.SetCookie(newCookie)

			data.IsLoggedIn = true
		}
	}

	return c.Render(http.StatusOK, "app", data)
}

func LoginForm(c echo.Context) error {
	return c.Render(http.StatusOK, "login-form", nil)
}

func RegisterForm(c echo.Context) error {
	return c.Render(http.StatusOK, "register-form", nil)
}

func Login(c echo.Context) error {
	token, err := panopticon.LoginCred(c.FormValue("handle"), c.FormValue("password"))
	if err != nil {
		e := ErrorData{
			Message: err.Error(),
		}

		return c.Render(http.StatusOK, "auth-screen", e)
	}

	newCookie := cookies.New(cookies.TokenCookieName, token)
	c.SetCookie(newCookie)

	return c.Render(http.StatusOK, "home-screen", nil)
}

func Register(c echo.Context) error {
	r := panopticon.Registration{
		Handle:          c.FormValue("handle"),
		DisplayName:     c.FormValue("displayName"),
		Password:        c.FormValue("password"),
		ConfirmPassword: c.FormValue("confirmPassword"),
	}

	err := panopticon.Register(r)
	if err != nil {
		e := ErrorData{
			Message: err.Error(),
		}

		return c.Render(http.StatusOK, "register-form", e)
	}

	return LoginForm(c)
}
