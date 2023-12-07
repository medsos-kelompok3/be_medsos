package routes

import (
	"be_medsos/features/comment"
	"be_medsos/features/posting"
	"be_medsos/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc user.Handler, pc posting.Handler, cc comment.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	RouteUser(e, uc)
	RoutePosting(e, pc)
	RouteComment(e, cc)
}

func RouteUser(e *echo.Echo, uc user.Handler) {
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())
	e.GET("/user/:username", uc.GetAllUserByUsername())
	e.DELETE("/user/:id", uc.Delete(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.PUT("/user/:id", uc.Update(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}

func RoutePosting(e *echo.Echo, pc posting.Handler) {
	e.POST("/posting", pc.Add(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/posting", pc.GetAll())
	e.PUT("/posting/:id", pc.Update(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.DELETE("/posting/:id", pc.Delete(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}

func RouteComment(e *echo.Echo, cc comment.Handler) {
	e.POST("/comment", cc.Add(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.PUT("/comment/:id", cc.Update(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
