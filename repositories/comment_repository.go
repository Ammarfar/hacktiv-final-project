package repositories

import (
	"finalproject/models"
	"finalproject/requests"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(request models.Comment) error
	List(userId any) ([]models.Comment, error)
	IsCommentOwner(userId any, commentId string) bool
	Update(request requests.CommentUpdateRequest) (*models.Comment, error)
	Delete(id any) error
}

type commentRepositoryImpl struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepositoryImpl{
		db: db,
	}
}

func (cr *commentRepositoryImpl) Create(request models.Comment) error {
	return cr.db.Create(&request).Error
}

func (cr *commentRepositoryImpl) List(userId any) ([]models.Comment, error) {
	var comments []models.Comment

	if err := cr.db.
		Preload("User").
		Preload("Photo").
		Where("user_id = ?", userId).
		Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (cr *commentRepositoryImpl) IsCommentOwner(userId any, commentId string) bool {
	return cr.db.Where("id = ?", commentId).Where("user_id = ?", userId).First(&models.Comment{}).Error == nil
}

func (cr *commentRepositoryImpl) Update(request requests.CommentUpdateRequest) (*models.Comment, error) {
	var comment models.Comment

	if err := cr.db.Where("id = ?", request.ID).Take(&comment).Error; err != nil {
		return nil, err
	}

	comment.Message = request.Message

	if err := cr.db.Save(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (cr *commentRepositoryImpl) Delete(id any) error {
	return cr.db.Delete(&models.Comment{}, id).Error
}
