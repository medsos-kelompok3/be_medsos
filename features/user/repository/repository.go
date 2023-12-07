package repository

import (
	"be_medsos/features/user"
	"errors"
	"strings"

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

type PostingModel struct {
	gorm.Model
	Caption       string
	GambarPosting string
	UserName      string
	User_id       uint
	Avatar        string
}

type Posting struct {
	ID            uint
	Caption       string
	GambarPosting string
	UserName      string
	Avatar        string
	JumlahKomen   int
}

type CommentModel struct {
	gorm.Model
	PostingID  uint
	IsiComment string
	UserName   string
	UserID     uint
	Avatar     string
}

type Comment struct {
	ID         uint
	PostingID  uint
	UserID     uint
	UserName   string
	Avatar     string
	IsiComment string
	CreatedAt  string
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
		err := errors.New("user tidak ditemukan")
		return nil, err
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

// ngambil info user dan profil
func (uq *UserQuery) GetProfil(id uint) (user.User, []Posting, error) {
	//ngambil user
	var userproses = new(user.User)
	userproses, err := uq.GetUserByID(id)
	if err != nil {
		return user.User{}, nil, err
	}
	// ngambil postingan
	var postingproses = new([]PostingModel)
	if err := uq.db.Find(&postingproses).Where("user_id = ?", id); err.Error != nil {
		if strings.Contains(err.Error.Error(), "not found") {
			errors.New("User tidak memiliki postingan, 404")

			return *userproses, nil, err.Error
		}
	}
	//ngambil jumlah komen
	var jumlahkomen []int
	for _, post := range *postingproses {
		var comments CommentModel
		var count int64
		uq.db.Model(&comments).Where("postingid = ?", post.ID).Count(&count)
		jumlahkomen = append(jumlahkomen, int(count))
	}

	// iterasi ke posting
	var postResponse = new([]Posting)
	for n, post := range *postingproses {
		isiposting := Posting{
			ID:            post.ID,
			Caption:       post.Caption,
			GambarPosting: post.GambarPosting,
			UserName:      post.UserName,
			Avatar:        post.Avatar,
			JumlahKomen:   jumlahkomen[n],
		}
		*postResponse = append(*postResponse, isiposting)
	}

	return *userproses, *postResponse, nil

}
