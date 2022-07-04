package repositories

import (
	"finalproject/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(request models.Photo) error
	List(userId any) ([]models.Photo, error)
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

func (pr *photoRepositoryImpl) List(userId any) ([]models.Photo, error) {
	var photos []models.Photo

	if err := pr.db.
		Preload("User").
		Where("user_id = ?", userId).
		Find(&photos).Error; err != nil {
		return nil, err
	}

	return photos, nil
}
