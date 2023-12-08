package comment

import (
	"be_medsos/features/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// type Comment struct {
// 	ID         uint
// 	PostingID  uint
// 	UserID     uint
// 	UserName   string
// 	Avatar     string
// 	IsiComment string
// 	CreatedAt  string
// }

type Handler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type Service interface {
	AddComment(token *jwt.Token, newComment models.Comment) (models.Comment, error)
	UpdateComment(token *jwt.Token, input models.Comment) (models.Comment, error)
	HapusComment(token *jwt.Token, commentID uint) error
}

type Repository interface {
	InsertComment(userID uint, postingID uint, newComment models.Comment) (models.Comment, error)
	UpdateComment(input models.Comment) (models.Comment, error)
	DeleteComment(commentID uint) error
}
