package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pezdel/go-auth/domain"
	"github.com/pezdel/go-auth/templates/partials"
)

func Register(users map[string]domain.User) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user domain.User
		err := c.Bind(&user)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		if user.Password != user.Confirm {
			return HTML(c, partials.PasswordMismatch(user.Email))
		}

		//save to db
		users[user.Email] = domain.User{
			Email:    user.Email,
			Password: user.Password,
		}
		c.Response().Header().Add("HX-Redirect", "/auth/login")

		return c.String(http.StatusOK, "")
	}
}

func CheckEmail(users map[string]domain.User) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user domain.User
		err := c.Bind(&user)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		_, taken := users[user.Email]
		if taken {
			return HTML(c, partials.EmailTaken())
		}

		return HTML(c, partials.ValidEmail(user.Email))
	}
}
