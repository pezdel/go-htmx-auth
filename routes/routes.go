package routes

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Serve(port string) {
	e := echo.New()
	e.HTTPErrorHandler = r.httpErrorHandler

	assetHandler := http.FileServer(rice.MustFindBox("../assets/").HTTPBox())
	e.GET("/", echo.WrapHandler(assetHandler))
	e.GET("/assets*", echo.WrapHandler(http.StripPrefix("/assets", assetHandler)))

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*.a.run.app", "*.vercel.app", "*://localhost:*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", r.HomePage)
	e.GET("/data", r.GetData)

	e.Logger.Fatal(e.Start(port))
}
