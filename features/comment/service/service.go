package service

import (
	"be_medsos/features/comment"
	"be_medsos/helper/jwt"
	"errors"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type CommentService struct {
	c comment.Repository
}

// HapusComment implements comment.Service.
func (cs *CommentService) HapusComment(token *golangjwt.Token, commentID uint) error {
	err := cs.c.DeleteComment(commentID)
	if err != nil {
		return errors.New("failed to delete the posting")
	}

	return nil
}

func New(model comment.Repository) comment.Service {
	return &CommentService{
		c: model,
	}
}

// AddComment implements comment.Service.
func (cs *CommentService) AddComment(token *golangjwt.Token, newComment comment.Comment) (comment.Comment, error) {
	userId, err := jwt.ExtractToken(token)
	if err != nil {
		return comment.Comment{}, err
	}
	result, err := cs.c.InsertComment(userId, newComment.PostingID, newComment)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return comment.Comment{}, errors.New("dobel input")
		}
		return comment.Comment{}, errors.New("terjadi kesalahan")
	}
	return result, nil
}

// UpdateComment implements comment.Service.
func (cs *CommentService) UpdateComment(token *golangjwt.Token, input comment.Comment) (comment.Comment, error) {
	respons, err := cs.c.UpdateComment(input)
	if err != nil {

		return comment.Comment{}, errors.New("kesalahan pada database")
	}
	return respons, nil
}
