package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pezdel/go-auth/templates/partials"
)

func Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "AuthJWT",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	return HTML(c, partials.NotLogged())
}

func AdminLogout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "AuthJWT",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	c.Response().Header().Add("HX-Redirect", "/")
	return c.String(http.StatusOK, "")
}
