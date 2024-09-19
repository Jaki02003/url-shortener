package routes

import (
	"fmt"

	"go-redis-url-shortener/controllers"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	echo          *echo.Echo
	urlController *controllers.UrlController
}

func New(echo *echo.Echo,
	urlController *controllers.UrlController,
) *Routes {
	return &Routes{
		echo:          echo,
		urlController: urlController,
	}
}

func (r *Routes) Init() {
	e := r.echo
	r.registerRoutes(e)
}

func (r *Routes) registerRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.POST("/encode", r.urlController.CreateShortUrl)
	v1.GET("/decode/:short-url", r.urlController.ReturnLongUrl)
	v1.GET("/:short-url", r.urlController.RedirectToActualUrl)

	for _, route := range e.Routes() {
		fmt.Println(fmt.Sprintf("%v || %v || %v", route.Method, route.Path, route.Name))
	}
}
