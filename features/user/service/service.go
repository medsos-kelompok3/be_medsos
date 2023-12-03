package service

import (
	"be_medsos/features/user"
	"be_medsos/helper/enkrip"
	"errors"
	"strings"
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
