package repository

import (
	"be_medsos/features/posting"
	pu "be_medsos/features/user"

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
func (pq *PostingQuery) InsertPosting(userName string, newPosting posting.Posting) (posting.Posting, error) {
	var inputDB = new(PostingModel)
	inputDB.Caption = newPosting.Caption
	inputDB.GambarPosting = newPosting.GambarPosting

	var user pu.User
	if err := pq.db.Where("Username = ?", userName).First(&user).Error; err != nil {
		// Handle error
		return posting.Posting{}, err
	}
	newPosting.UserName = user.Username
	// Simpan posting ke dalam database
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
