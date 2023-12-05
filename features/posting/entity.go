package posting

import (
	"github.com/golang-jwt/jwt"
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
}

type Service interface {
	AddPosting(token *jwt.Token, newPosting Posting) (Posting, error)
}

type Repository interface {
	InsertPosting(userName string, newPosting Posting) (Posting, error)
}
