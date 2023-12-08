package posting

import (
	"be_medsos/features/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// type Posting struct {
// 	PostingID            uint
// 	Caption       string
// 	GambarPosting string
// 	UserName      string
// }

type Handler interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetOne() echo.HandlerFunc
}

type Service interface {
	AddPosting(token *jwt.Token, newPosting models.Posting) (models.Posting, error)
	UpdatePosting(token *jwt.Token, input models.Posting) (models.Posting, error)
	SemuaPosting(page, limit int) ([]models.Posting, error)
	HapusPosting(token *jwt.Token, postingID uint) error
	GetOne(id uint) (models.Posting, []models.Comment, error)
}

type Repository interface {
	InsertPosting(userID uint, newPosting models.Posting) (models.Posting, error)
	DeletePosting(postingID uint) error
	UpdatePosting(input models.Posting) (models.Posting, error)
	GetTanpaPosting(page, limit int) ([]models.Posting, error)
	GetOne(id uint) (models.Posting, []models.Comment, error)
}
