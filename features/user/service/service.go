package service

import (
	"be_medsos/features/user"
	"be_medsos/helper/enkrip"
	"be_medsos/helper/jwt"
	"errors"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	repo user.Repository
	h    enkrip.HashInterface
}

func New(r user.Repository, h enkrip.HashInterface) user.Service {
	return &UserService{
		repo: r,
		h:    h,
	}
}

func (us *UserService) AddUser(input user.User) error {
	if input.Username == "" || input.Password == "" {
		return errors.New("username and password are required")
	}
	ePassword, err := us.h.HashPassword(input.Password)

	if err != nil {
		return errors.New("terjadi error saat enkripsi")
	}
	input.Password = ePassword

	result := us.repo.AddUser(input)
	if result != nil {
		if strings.Contains(err.Error(), "didaftarkan") {
			return result
		}
		return errors.New("terjadi kesalahan pada sistem")
	}
	return result
}

// Login implements user.Service.
func (us *UserService) Login(username string, password string) (user.User, error) {
	if username == "" || password == "" {
		return user.User{}, errors.New("username and password are required")
	}
	result, err := us.repo.Login(username)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return user.User{}, errors.New("data tidak ditemukan")
		}
		return user.User{}, errors.New("terjadi kesalahan pada sistem")
	}

	err = us.h.Compare(result.Password, password)

	if err != nil {
		return user.User{}, errors.New("password salah")
	}

	return result, nil
}

// DapatUser implements user.Service.
func (us *UserService) DapatUser(username string) (user.User, error) {
	result, err := us.repo.GetUserByUsername(username)
	if err != nil {
		return user.User{}, errors.New("failed to retrieve inserted Data")
	}
	return result, nil
}

// HapusUser implements user.Service.
func (us *UserService) HapusUser(token *golangjwt.Token, userID uint) error {
	userId, err := jwt.ExtractToken(token)
	if err != nil {
		return err
	}
	exitingUser, err := us.repo.GetUserByID(userID)
	if err != nil {
		return errors.New("failed to retrieve the user for deletion")
	}
	if exitingUser.ID != userId {
		return errors.New("you don't have permission to delete this user")
	}
	err = us.repo.DeleteUser(userID)
	if err != nil {
		return errors.New("failed to delete the user")
	}

	return nil
}

// update user
func (us *UserService) UpdateUser(token *golangjwt.Token, input user.User) (user.User, error) {
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return user.User{}, errors.New("harap login")
	}
	if userID != input.ID {

		return user.User{}, errors.New("id tidak cocok")
	}
	// ambil data user yg lama
	base, err := us.repo.GetUserByID(userID)
	if err != nil {
		return user.User{}, errors.New("user tidak ditemukan")
	}
	if input.Password != "" {
		err = us.h.Compare(base.Password, input.Password)

		if err != nil {
			return user.User{}, errors.New("password salah")
		}
	}
	if input.NewPassword != "" {
		newpass, err := us.h.HashPassword(input.NewPassword)
		if err != nil {
			return user.User{}, errors.New("masukkan password baru dengan benar")
		}
		input.NewPassword = newpass
	}

	respons, err := us.repo.UpdateUser(input)
	if err != nil {

		return user.User{}, errors.New("kesalahan pada database")
	}
	return respons, nil

}
