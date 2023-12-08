package repository

import (
	"be_medsos/features/models"
	"be_medsos/features/posting"
	"fmt"

	"gorm.io/gorm"
)

// type PostingModel struct {
// 	gorm.Model
// 	Caption       string
// 	GambarPosting string
// 	UserName      string
// 	User_id       uint
// 	Avatar        uint
// }

// type CommentModel struct {
// 	gorm.Model
// 	PostingID  uint
// 	IsiComment string
// 	UserName   string
// 	UserID     uint
// 	Avatar     string
// }

type PostingQuery struct {
	db *gorm.DB
}

// DeletePosting implements posting.Repository.
func (pq *PostingQuery) DeletePosting(postingID uint) error {
	var postingModel models.PostingModel

	if err := pq.db.First(&postingModel, postingID).Error; err != nil {
		return err
	}
	if err := pq.db.Delete(&postingModel, postingID).Error; err != nil {
		return err
	}

	return nil
}

// UpdatePosting implements posting.Repository.
func (pq *PostingQuery) UpdatePosting(input models.Posting) (models.Posting, error) {
	var proses models.PostingModel
	if err := pq.db.First(&proses, input.ID).Error; err != nil {
		return models.Posting{}, err
	}

	proses.Caption = input.Caption
	proses.GambarPosting = input.GambarPosting

	if err := pq.db.Save(&proses).Error; err != nil {

		return models.Posting{}, err
	}
	result := models.Posting{
		ID:            proses.ID,
		Caption:       proses.Caption,
		GambarPosting: proses.GambarPosting,
		UserName:      proses.UserName,
	}

	return result, nil
}

// GetTanpaPosting implements posting.Repository.
func (pq *PostingQuery) GetTanpaPosting(page int, limit int) ([]models.Posting, error) {
	var postings []models.PostingModel
	offset := (page - 1) * limit
	if err := pq.db.Offset(offset).Limit(limit).Find(&postings).Error; err != nil {
		return nil, err
	}
	var result []models.Posting
	for _, s := range postings {
		result = append(result, models.Posting{
			ID:            s.ID,
			Caption:       s.Caption,
			GambarPosting: s.GambarPosting,
			UserName:      s.UserName,
		})
	}
	return result, nil
}

// InsertPosting implements posting.Repository.
func (pq *PostingQuery) InsertPosting(userID uint, newPosting models.Posting) (models.Posting, error) {
	var inputDB = new(models.PostingModel)
	inputDB.Caption = newPosting.Caption
	inputDB.GambarPosting = newPosting.GambarPosting

	var user models.UserModel
	if err := pq.db.First(&user, userID).Error; err != nil {
		fmt.Println("Error mengambil data customer:", err)
		return models.Posting{}, err
	}
	inputDB.UserName = user.Username
	inputDB.Avatar = user.Avatar
	inputDB.User_id = userID

	if err := pq.db.Create(&inputDB).Error; err != nil {
		// Handle error saat menyimpan posting ke dalam database
		return models.Posting{}, err
	}

	newPosting.ID = inputDB.ID
	newPosting.UserName = inputDB.UserName

	return newPosting, nil
}

func New(db *gorm.DB) posting.Repository {
	return &PostingQuery{
		db: db,
	}
}
