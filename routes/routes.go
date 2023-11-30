package routes

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"golang.org/x/time/rate"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pezdel/go-auth/domain"
	"github.com/pezdel/go-auth/routes/handlers"
	mw "github.com/pezdel/go-auth/routes/middleware"
)

var users = map[string]domain.User{
	"a.a@gmail.com": {Email: "a.a@gmail.com", Password: "123"},
}

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
		rate.Limit(10),
	)))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*.a.run.app", "*.vercel.app", "*://localhost:*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	//HomePage
	e.GET("/", handlers.HomePage)

	//AuthPages
	auth := e.Group("/auth")
	auth.Use(mw.RedirectIfAuthenticated)
	auth.GET("/login", handlers.LoginPage)
	auth.GET("/register", handlers.RegisterPage)
	auth.GET("/success", handlers.SuccessPage)

	//AdminPages
	admin := e.Group("/admin")
	admin.Use(mw.RedirectIfNotAuthenticated)
	admin.GET("", handlers.AdminPage)

	//handlers
	e.POST("/login", handlers.Login(users))
	e.POST("/register", handlers.Register(users))
	e.POST("/checkEmail", handlers.CheckEmail(users))
	e.GET("/logout", handlers.Logout)
	e.GET("/adminLogout", handlers.AdminLogout)

	e.Logger.Fatal(e.Start(port))
}
