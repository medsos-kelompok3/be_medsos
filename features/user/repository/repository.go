package repository

import (
	"be_medsos/features/user"
	"errors"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
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

// add new user
func (uq *UserQuery) AddUser(input user.User) error {
	var newUser = new(UserModel)
	newUser.Username = input.Username
	newUser.Email = input.Email
	newUser.Password = input.Password
	newUser.Address = input.Address
	if err := uq.db.Create(&newUser).Error; err != nil {
		return errors.New("username/email sudah didaftarkan")
	}

	return nil
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

func (uq *UserQuery) UpdateUser(input user.User) (user.User, error) {
	var proses UserModel
	if err := uq.db.First(&proses, input.ID).Error; err != nil {
		return user.User{}, err
	}

	// Jika tidak ada buku ditemukan
	if proses.ID == 0 {
		err := errors.New("user tidak ditemukan")
		return user.User{}, err
	}

	if input.Username != "" {
		proses.Username = input.Username
	}
	if input.Email != "" {
		proses.Email = input.Email
	}
	if input.Bio != "" {
		proses.Bio = input.Bio
	}
	if input.Avatar != "" {
		proses.Avatar = input.Avatar
	}

	if input.Address != "" {
		proses.Address = input.Address
	}
	if input.NewPassword != "" {
		proses.Password = input.NewPassword
	}

	if err := uq.db.Save(&proses).Error; err != nil {

		return user.User{}, err
	}
	result := user.User{
		ID:       proses.ID,
		Username: proses.Username,
		Email:    proses.Email,
		Address:  proses.Address,
		Avatar:   proses.Avatar,
		Password: proses.Password,
		Bio:      proses.Bio,
	}

	return result, nil
}
