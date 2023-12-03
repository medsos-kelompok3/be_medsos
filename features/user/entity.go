package user

import "github.com/labstack/echo/v4"

type User struct {
	ID       uint
	Username string
	Email    string
	Address  string
	Bio      string
	Avatar   string
	Password string
}

type Handler interface {
	Login() echo.HandlerFunc
}

type Service interface {
	Login(username string, password string) (User, error)
}

type Repository interface {
	Login(username string) (User, error)
}
