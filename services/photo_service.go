package services

import (
	"finalproject/models"
	"finalproject/repositories"

	"gorm.io/gorm"
)

type PhotoService interface {
	Create(request models.Photo) error
	List(userId any) ([]models.Photo, error)
}

type photoServiceImpl struct {
	repository repositories.PhotoRepository
}

func NewPhotoService(db *gorm.DB) PhotoService {
	return &photoServiceImpl{
		repository: repositories.NewPhotoRepository(db),
	}
}

func (pc *photoServiceImpl) Create(request models.Photo) error {
	return pc.repository.Create(request)
}

func (pc *photoServiceImpl) List(userId any) ([]models.Photo, error) {
	return pc.repository.List(userId)
}
