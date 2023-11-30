package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// middleware for admin pages
func RedirectIfNotAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie("AuthJWT")
		if err != nil {
			// User is not logged in, redirect to login page
			return c.Redirect(http.StatusFound, "/auth/login")
		}

		// User is logged in, continue to the next handler
		return next(c)
	}
}

// middleware for Auth pages
func RedirectIfAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie("AuthJWT")
		if err == nil {
			return c.Redirect(http.StatusFound, "/admin")
		}
		return next(c)
	}
}
