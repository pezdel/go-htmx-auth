package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pezdel/go-auth/domain"
	"github.com/pezdel/go-auth/templates/partials"
)

func Login(users map[string]domain.User) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user domain.User
		err := c.Bind(&user)
		if err != nil {
			return HTML(c, partials.LoginInvalid(user.Email))
		}

		storedUser, ok := users[user.Email]
		if !ok || storedUser.Password != user.Password {
			return HTML(c, partials.LoginInvalid(user.Email))
		}

		//valid user

		c.SetCookie(createCookie("We created some cookie"))
		c.Response().Header().Add("HX-Redirect", "/admin")
		return c.String(http.StatusOK, "")
	}
}

func createCookie(name string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "AuthJWT",
		Value:    "someValue",
		Expires:  time.Now().Add(24 * time.Hour), // Expires in 24 hours
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true, // Set to true if serving over HTTPS
		// Add other cookie options as needed
	}

	return cookie
}
