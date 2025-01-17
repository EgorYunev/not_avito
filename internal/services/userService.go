package services

import (
	"crypto/sha1"
	"fmt"

	"github.com/EgorYunev/not_avito/internal/data"
	"github.com/EgorYunev/not_avito/internal/models"
)

var salt = "hguefkj42dskalf3kf1"

type UserService struct {
	UserRepository *data.UserRepository
}

func (s *UserService) CreateUser(user *models.User) error {

	user.Password = generateHashPassword(user.Password)
	fmt.Println(user.Password)

	err := s.UserRepository.Insert(user)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Authorize(email, password string) (bool, error) {

	user, err := s.UserRepository.GetByEmail(email)

	if err != nil {
		return false, err
	}

	pass := generateHashPassword(password)

	if pass == user.Password {
		return true, nil
	} else {
		return false, nil
	}

}

func (s *UserService) GetById(id int) (*models.User, error) {

	user, err := s.UserRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func generateHashPassword(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
