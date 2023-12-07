package repository

import (
	"be_medsos/features/posting"
	pr "be_medsos/features/user/repository"
	"fmt"

	"gorm.io/gorm"
)

type PostingModel struct {
	gorm.Model
	Caption       string
	GambarPosting string
	UserName      string
	UserID        uint
}

type PostingQuery struct {
	db *gorm.DB
}

// UpdatePosting implements posting.Repository.
func (pq *PostingQuery) UpdatePosting(input posting.Posting) (posting.Posting, error) {
	var proses PostingModel
	if err := pq.db.First(&proses, input.ID).Error; err != nil {
		return posting.Posting{}, err
	}

	proses.Caption = input.Caption
	proses.GambarPosting = input.GambarPosting

	if err := pq.db.Save(&proses).Error; err != nil {

		return posting.Posting{}, err
	}
	result := posting.Posting{
		ID:            proses.ID,
		Caption:       proses.Caption,
		GambarPosting: proses.GambarPosting,
		UserName:      proses.UserName,
	}

	return result, nil
}

// GetTanpaPosting implements posting.Repository.
func (pq *PostingQuery) GetTanpaPosting(page int, limit int) ([]posting.Posting, error) {
	var postings []PostingModel
	offset := (page - 1) * limit
	if err := pq.db.Offset(offset).Limit(limit).Find(&postings).Error; err != nil {
		return nil, err
	}
	var result []posting.Posting
	for _, s := range postings {
		result = append(result, posting.Posting{
			ID:            s.ID,
			Caption:       s.Caption,
			GambarPosting: s.GambarPosting,
			UserName:      s.UserName,
		})
	}
	return result, nil
}

// InsertPosting implements posting.Repository.
func (pq *PostingQuery) InsertPosting(userID uint, newPosting posting.Posting) (posting.Posting, error) {
	var inputDB = new(PostingModel)
	inputDB.Caption = newPosting.Caption
	inputDB.GambarPosting = newPosting.GambarPosting

	var user pr.UserModel
	if err := pq.db.First(&user, userID).Error; err != nil {
		fmt.Println("Error mengambil data customer:", err)
		return posting.Posting{}, err
	}
	inputDB.UserName = user.Username

	if err := pq.db.Create(&inputDB).Error; err != nil {
		// Handle error saat menyimpan posting ke dalam database
		return posting.Posting{}, err
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
