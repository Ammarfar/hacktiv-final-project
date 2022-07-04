package repositories

import (
	"finalproject/models"
	"finalproject/requests"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(request models.Photo) error
	List(userId any) ([]models.Photo, error)
	IsPhotoOwner(userId any, photoId string) bool
	Update(request requests.PhotoUpdateRequest) (*models.Photo, error)
	Delete(id any) error
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

func (pr *photoRepositoryImpl) IsPhotoOwner(userId any, photoId string) bool {
	return pr.db.Where("id = ?", photoId).Where("user_id = ?", userId).First(&models.Photo{}).Error == nil
}

func (pr *photoRepositoryImpl) Update(request requests.PhotoUpdateRequest) (*models.Photo, error) {
	var photo models.Photo

	if err := pr.db.Where("id = ?", request.ID).Take(&photo).Error; err != nil {
		return nil, err
	}

	photo.Title = request.Title
	photo.Caption = request.Caption
	photo.PhotoUrl = request.PhotoUrl

	if err := pr.db.Save(&photo).Error; err != nil {
		return nil, err
	}

	return &photo, nil
}

func (pr *photoRepositoryImpl) Delete(id any) error {
	return pr.db.Delete(&models.Photo{}, id).Error
}
