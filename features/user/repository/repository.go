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

// GetUserByUsername implements user.Repository.
func (uq *UserQuery) GetUserByUsername(username string) (user.User, error) {
	var userModel UserModel
	if err := uq.db.Where("username = ?", username).First(&userModel).Error; err != nil {
		return user.User{}, err
	}

	result := user.User{
		ID:       userModel.ID,
		Username: userModel.Username,
		Bio:      userModel.Bio,
		Avatar:   userModel.Avatar,
	}

	return result, nil
}

// DeleteUser implements user.Repository.
func (uq *UserQuery) DeleteUser(userID uint) error {
	var exitingUser UserModel

	if err := uq.db.First(&exitingUser, userID).Error; err != nil {
		return err
	}

	if err := uq.db.Delete(&exitingUser).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByID implements user.Repository.
func (uq *UserQuery) GetUserByID(userID uint) (*user.User, error) {
	var userModel UserModel
	if err := uq.db.First(&userModel, userID).Error; err != nil {
		return nil, err
	}

	// Jika tidak ada buku ditemukan
	if userModel.ID == 0 {
		return nil, nil
	}

	result := &user.User{
		ID:       userModel.ID,
		Username: userModel.Username,
		Email:    userModel.Email,
		Address:  userModel.Address,
		Avatar:   userModel.Avatar,
		Password: userModel.Password,
	}

	return result, nil
}
