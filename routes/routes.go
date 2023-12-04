package routes

import (
	"be_medsos/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc user.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	RouteUser(e, uc)
}

func RouteUser(e *echo.Echo, uc user.Handler) {
	e.POST("/login", uc.Login())
	e.GET("/user/:username", uc.GetAllUserByUsername())
	e.DELETE("/user/:id", uc.Delete(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
