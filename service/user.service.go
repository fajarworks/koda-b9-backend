package service

import (
	"errors"

	"github.com/fajarworks/koda-b9-backend/dto"
	// "github.com/fajarworks/koda-b9-backend/model"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

var Users = []dto.User{}

func (s *UserService) WrongEmailAndPass(form dto.User) error {
	for _, v := range Users {
		if form.Email != v.Email || form.Password != v.Password {

			return errors.New("email atau password salah")
		}

	}
	return nil
}

func (s *UserService) EmailExist(form dto.User) error {
	for _, v := range Users {
		if v.Email == form.Email {

			return errors.New("email sudah tersedia")
		}

	}
	return nil
}

func (s *UserService) LengthPassword(form dto.User) error {
	if len(form.Password) < 8 {
		return errors.New("password harus lebih dari 8 karakter")

	}
	return nil
}
