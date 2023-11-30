package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pezdel/go-auth/templates/pages"
)

func HomePage(c echo.Context) error {
	_, err := c.Cookie("AuthJWT")
	if err != nil {
		return HTML(c, pages.HomePage(false))
	}
	return HTML(c, pages.HomePage(true))
}

func LoginPage(c echo.Context) error {
	return HTML(c, pages.LoginPage())
}

func RegisterPage(c echo.Context) error {
	return HTML(c, pages.RegisterPage())
}

func SuccessPage(c echo.Context) error {
	return HTML(c, pages.SuccessPage())
}

func AdminPage(c echo.Context) error {
	val, _ := c.Cookie("AuthJWT")
	return HTML(c, pages.AdminPage(val.Value))
}

func ErrorPage(c echo.Context) error {
	return HTML(c, pages.ErrorPage())
}

func HTML(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
