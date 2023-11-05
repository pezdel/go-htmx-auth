package routes

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pezdel/go-template/templates"
)

func (r *Router) HomePage(c echo.Context) error {
	return HTML(c, templates.Home("Home Page"))
}

func (r *Router) ErrorPage(c echo.Context) error {
	return HTML(c, templates.ErrorPage())
}

func HTML(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
