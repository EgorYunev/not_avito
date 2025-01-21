package services

import (
	"github.com/EgorYunev/not_avito/internal/data"
	"github.com/EgorYunev/not_avito/internal/models"
)

type AdService struct {
	Repo *data.AdRepository
}

func (s *AdService) CreateAd(ad *models.Ad) error {

	err := s.Repo.Insert(ad)

	if err != nil {
		return err
	}

	return nil

}

func (s *AdService) Delete(adId int, email string) error {

	err := s.Repo.Delete(adId, email)

	if err != nil {
		return err
	}

	return nil
}

func (s *AdService) ChangeAd(ad *models.Ad, email string) error {

	err := s.Repo.ChangeAd(ad, email)
	if err != nil {
		return err
	}

	return nil
}
