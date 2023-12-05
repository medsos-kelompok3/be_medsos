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
}

type PostingQuery struct {
	db *gorm.DB
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
