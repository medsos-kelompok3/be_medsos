package comment

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Comment struct {
	ID         uint
	PostingID  uint
	UserID     uint
	UserName   string
	IsiComment string
}

type Handler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type Service interface {
	AddComment(token *jwt.Token, newComment Comment) (Comment, error)
	UpdateComment(token *jwt.Token, input Comment) (Comment, error)
	HapusComment(token *jwt.Token, commentID uint) error
}

type Repository interface {
	InsertComment(userID uint, postingID uint, newComment Comment) (Comment, error)
	UpdateComment(input Comment) (Comment, error)
	DeleteComment(commentID uint) error
}
