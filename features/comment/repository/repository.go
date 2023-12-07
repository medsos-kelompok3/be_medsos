package repository

import (
	"be_medsos/features/comment"
	cr "be_medsos/features/posting/repository"
	pr "be_medsos/features/user/repository"
	"fmt"

	"gorm.io/gorm"
)

type CommentModel struct {
	gorm.Model
	PostingID  uint
	IsiComment string
	UserName   string
}

type CommentQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.Repository {
	return &CommentQuery{
		db: db,
	}
}

// InsertComment implements comment.Repository.
func (cq *CommentQuery) InsertComment(userID uint, postingID uint, newComment comment.Comment) (comment.Comment, error) {
	var inputDB = new(CommentModel)
	inputDB.IsiComment = newComment.IsiComment

	var user pr.UserModel
	if err := cq.db.First(&user, userID).Error; err != nil {
		fmt.Println("Error mengambil data customer:", err)
		return comment.Comment{}, err
	}
	inputDB.UserName = user.Username

	var posting cr.PostingModel
	if err := cq.db.First(&posting, postingID).Error; err != nil {
		fmt.Println("Error mengambil data posting:", err)
		return comment.Comment{}, err
	}
	inputDB.PostingID = posting.ID

	if err := cq.db.Create(&inputDB).Error; err != nil {
		// Handle error saat menyimpan posting ke dalam database
		return comment.Comment{}, err
	}

	newComment.ID = inputDB.ID
	newComment.UserName = inputDB.UserName
	newComment.PostingID = inputDB.PostingID

	return newComment, nil
}
