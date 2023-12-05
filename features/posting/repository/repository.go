package repository

import (
	"be_medsos/features/posting"
	pu "be_medsos/features/user"
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

	var user pu.User
	if err := pq.db.First(&user, user.ID).Error; err != nil {
		fmt.Println("Error mengambil data customer:", err)
		return posting.Posting{}, err
	}
	inputDB.UserName = user.Username

	if err := pq.db.Create(&newPosting).Error; err != nil {
		// Handle error saat menyimpan posting ke dalam database
		return posting.Posting{}, err
	}

	return newPosting, nil
}

func New(db *gorm.DB) posting.Repository {
	return &PostingQuery{
		db: db,
	}
}
