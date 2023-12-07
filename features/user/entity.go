package user

import (
	posting "be_medsos/features/posting"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID          uint
	Username    string
	Email       string
	Address     string
	Bio         string
	Avatar      string
	Password    string
	NewPassword string
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetAllUserByUsername() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetUserDetails() echo.HandlerFunc
	GetUserProfiles() echo.HandlerFunc
}

type Service interface {
	AddUser(input User) error
	Login(username string, password string) (User, error)
	DapatUser(username string) (User, error)
	HapusUser(token *jwt.Token, userID uint) error
	UpdateUser(token *jwt.Token, input User) (User, error)
	GetUserDetails(token *jwt.Token, userID uint) (User, error)
	GetUserProfiles(token *jwt.Token, userID uint) (User, []posting.Posting, error)
}

type Repository interface {
	AddUser(input User) error
	Login(username string) (User, error)
	GetUserByUsername(username string) (User, error)
	GetUserByID(userID uint) (*User, error)
	DeleteUser(userID uint) error
	UpdateUser(input User) (User, error)
	GetProfil(id uint) (User, []posting.Posting, error)
}
