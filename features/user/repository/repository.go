package repository

import (
	"be_medsos/features/user"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username string
	Email    string
	Address  string
	Bio      string
	Avatar   string
	Password string
}

type UserQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &UserQuery{
		db: db,
	}
}

// Login implements user.Repository.
func (uq *UserQuery) Login(username string) (user.User, error) {
	var userData = new(UserModel)

	if err := uq.db.Where("username = ?", username).First(userData).Error; err != nil {
		return user.User{}, err
	}

	var result = new(user.User)
	result.ID = userData.ID
	result.Username = userData.Username
	result.Password = userData.Password

	return *result, nil
}
