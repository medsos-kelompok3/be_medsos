package user

import (
	"github.com/labstack/echo/v4"
)

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
	GetAllUserByUsername() echo.HandlerFunc
}

type Service interface {
	Login(username string, password string) (User, error)
	DapatUser(username string) (User, error)
}

type Repository interface {
	Login(username string) (User, error)
	GetUserByUsername(username string) (User, error)
}
