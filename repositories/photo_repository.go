package repositories

import (
	"finalproject/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(request models.Photo) error
}

type photoRepositoryImpl struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepositoryImpl{
		db: db,
	}
}

func (pr *photoRepositoryImpl) Create(request models.Photo) error {
	return pr.db.Create(&request).Error
}
