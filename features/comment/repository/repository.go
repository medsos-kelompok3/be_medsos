package repository

import (
	"be_medsos/features/comment"

	"be_medsos/features/models"
	"fmt"

	"gorm.io/gorm"
)

// type CommentModel struct {
// 	gorm.Model
// 	PostingID  uint
// 	IsiComment string
// 	UserName   string
// 	UserID     uint
// 	Avatar     string
// }

type CommentQuery struct {
	db *gorm.DB
}

// DeleteComment implements comment.Repository.
func (cq *CommentQuery) DeleteComment(commentID uint) error {
	var commentModel models.CommentModel

	if err := cq.db.First(&commentModel, commentID).Error; err != nil {
		return err
	}
	if err := cq.db.Delete(&commentModel, commentID).Error; err != nil {
		return err
	}

	return nil
}

func New(db *gorm.DB) comment.Repository {
	return &CommentQuery{
		db: db,
	}
}

// InsertComment implements comment.Repository.
func (cq *CommentQuery) InsertComment(userID uint, postingID uint, newComment models.Comment) (models.Comment, error) {
	var inputDB = new(models.CommentModel)
	inputDB.IsiComment = newComment.IsiComment

	var user models.UserModel
	if err := cq.db.First(&user, userID).Error; err != nil {
		fmt.Println("Error mengambil data customer:", err)
		return models.Comment{}, err
	}
	inputDB.UserName = user.Username
	inputDB.UserID = user.ID
	inputDB.Avatar = user.Avatar

	var posting models.PostingModel
	if err := cq.db.First(&posting, postingID).Error; err != nil {
		fmt.Println("Error mengambil data posting:", err)
		return models.Comment{}, err
	}
	inputDB.PostingID = posting.ID

	if err := cq.db.Create(&inputDB).Error; err != nil {
		// Handle error saat menyimpan posting ke dalam database
		return models.Comment{}, err
	}

	newComment.ID = inputDB.ID
	newComment.UserName = inputDB.UserName
	newComment.PostingID = inputDB.PostingID

	return newComment, nil
}

// UpdateComment implements comment.Repository.
func (cq *CommentQuery) UpdateComment(input models.Comment) (models.Comment, error) {
	var proses models.CommentModel
	if err := cq.db.First(&proses, input.ID).Error; err != nil {
		return models.Comment{}, err
	}

	proses.IsiComment = input.IsiComment

	if err := cq.db.Save(&proses).Error; err != nil {

		return models.Comment{}, err
	}
	result := models.Comment{
		ID:         proses.ID,
		IsiComment: proses.IsiComment,
		PostingID:  proses.PostingID,
		UserName:   proses.UserName,
	}

	return result, nil
}

func (cq *CommentQuery) GetOne(id uint) (models.Comment, error) {
	var proses models.CommentModel
	if err := cq.db.First(&proses).Where("id = ?", id).Error; err != nil {

		return models.Comment{}, err
	}
	hasil := models.Comment{
		ID:         proses.ID,
		Avatar:     proses.Avatar,
		IsiComment: proses.IsiComment,
		UserName:   proses.UserName,
		CreatedAt:  proses.CreatedAt.String(),
		UserID:     proses.UserID,
		PostingID:  proses.PostingID,
	}
	return hasil, nil
}
