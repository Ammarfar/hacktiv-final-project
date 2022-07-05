package repositories

import (
	"finalproject/models"
	"finalproject/requests"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(request models.SocialMedia) error
	List(userId any) ([]models.SocialMedia, error)
	IsSocialMediaOwner(userId any, socialMediaId string) bool
	Update(request requests.SocialMediaUpdateRequest) (*models.SocialMedia, error)
	Delete(id any) error
}

type socialMediaRepositoryImpl struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepositoryImpl{
		db: db,
	}
}

func (pr *socialMediaRepositoryImpl) Create(request models.SocialMedia) error {
	return pr.db.Create(&request).Error
}

func (pr *socialMediaRepositoryImpl) List(userId any) ([]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia

	if err := pr.db.
		Preload("User").
		Where("user_id = ?", userId).
		Find(&socialMedias).Error; err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func (pr *socialMediaRepositoryImpl) IsSocialMediaOwner(userId any, socialMediaId string) bool {
	return pr.db.Where("id = ?", socialMediaId).Where("user_id = ?", userId).First(&models.SocialMedia{}).Error == nil
}

func (pr *socialMediaRepositoryImpl) Update(request requests.SocialMediaUpdateRequest) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia

	if err := pr.db.Where("id = ?", request.ID).Take(&socialMedia).Error; err != nil {
		return nil, err
	}

	socialMedia.Name = request.Name
	socialMedia.SocialMediaUrl = request.SocialMediaUrl

	if err := pr.db.Save(&socialMedia).Error; err != nil {
		return nil, err
	}

	return &socialMedia, nil
}

func (pr *socialMediaRepositoryImpl) Delete(id any) error {
	return pr.db.Delete(&models.SocialMedia{}, id).Error
}
