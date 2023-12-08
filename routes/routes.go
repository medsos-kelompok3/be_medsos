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
	e.GET("/user/:username", uc.GetAllUserByUsername(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/user/:user_id", uc.GetUserDetails(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/user/:user_id/profile", uc.GetUserProfiles(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.DELETE("/user/:user_id", uc.Delete(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.PUT("/user/:user_id", uc.Update(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}

func RoutePosting(e *echo.Echo, pc posting.Handler) {
	e.POST("/posting", pc.Add(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/posting", pc.GetAll())
	e.GET("/posting/:posting_id", pc.GetOne()) //single post
	e.PUT("/posting/:posting_id", pc.Update(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.DELETE("/posting/:posting_id", pc.Delete(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}

func RouteComment(e *echo.Echo, cc comment.Handler) {
	e.POST("/comment", cc.Add(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/comment/:comment_id", cc.GetOne(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.PUT("/comment/:comment_id", cc.Update(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.DELETE("/comment/:comment_id", cc.Delete(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
