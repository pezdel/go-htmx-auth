package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) GetData(c echo.Context) error {
	return c.String(http.StatusOK, "sample api endpoint")
}
