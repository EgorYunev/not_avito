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

	err := s.UserRepository.Insert(user)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetById(id int) (*models.User, error) {

	user, err := s.GetById(id)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func generateHashPassword(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%s", hash.Sum([]byte(salt)))
}
