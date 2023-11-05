package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	InvalidAuthToken = "InvalidAuthToken"
	ValidationError  = "ValidationError"
	ServerError      = "ServerError"
)

type Response struct {
	Data      interface{} `json:"data"`
	Meta      interface{} `json:"meta,omitempty"`
	Message   string      `json:"message,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
}

func (r *Router) httpErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if code != http.StatusInternalServerError {
		_ = r.ErrorPage(c)
	} else {
		// log.Error(err)
		_ = HTTPServerError(c)
	}
}

func HTTPServerError(c echo.Context) error {
	r := Response{
		ErrorCode: ServerError,
		Message:   "Something went wrong",
	}

	return c.JSON(http.StatusInternalServerError, r)
}
