package posting

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Posting struct {
	ID            uint
	Caption       string
	GambarPosting string
	UserName      string
}

type Handler interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type Service interface {
	AddPosting(token *jwt.Token, newPosting Posting) (Posting, error)
	UpdatePosting(token *jwt.Token, input Posting) (Posting, error)
	SemuaPosting(page, limit int) ([]Posting, error)
}

type Repository interface {
	InsertPosting(userID uint, newPosting Posting) (Posting, error)
	UpdatePosting(input Posting) (Posting, error)
	GetTanpaPosting(page, limit int) ([]Posting, error)
}
