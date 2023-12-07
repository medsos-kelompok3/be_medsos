package comment

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Comment struct {
	ID         uint
	PostingID  uint
	UserName   string
	IsiComment string
}

type Handler interface {
	Add() echo.HandlerFunc
}

type Service interface {
	AddComment(token *jwt.Token, newComment Comment) (Comment, error)
}

type Repository interface {
	InsertComment(userID uint, postingID uint, newComment Comment) (Comment, error)
}
