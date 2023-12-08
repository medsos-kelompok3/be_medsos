package user

import (
	"be_medsos/features/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// type User struct {
// 	ID          uint
// 	Username    string
// 	Email       string
// 	Address     string
// 	Bio         string
// 	Avatar      string
// 	Password    string
// 	NewPassword string
// }

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
	AddUser(input models.User) error
	Login(username string, password string) (models.User, error)
	DapatUser(username string) (models.User, error)
	HapusUser(token *jwt.Token, userID uint) error
	UpdateUser(token *jwt.Token, input models.User) (models.User, error)
	GetUserDetails(token *jwt.Token, userID uint) (models.User, error)
	GetUserProfiles(token *jwt.Token, userID uint) (models.User, []models.Posting, error)
}

type Repository interface {
	AddUser(input models.User) error
	Login(username string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	GetUserByID(userID uint) (*models.User, error)
	DeleteUser(userID uint) error
	UpdateUser(input models.User) (models.User, error)
	GetProfil(id uint) (models.User, []models.Posting, error)
}
