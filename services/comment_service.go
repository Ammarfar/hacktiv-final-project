package services

import (
	"errors"
	"finalproject/configs"
	"finalproject/models"
	"finalproject/repositories"
	"finalproject/requests"
	"finalproject/responses"

	"gorm.io/gorm"
)

type CommentService interface {
	Create(request models.Comment) error
	List(userId any) ([]responses.CommentListResponse, error)
	Update(request requests.CommentUpdateRequest) (*models.Comment, error)
	Delete(userId any, commentId string) error
}

type commentServiceImpl struct {
	repository repositories.CommentRepository
}

func NewCommentService(db *gorm.DB) CommentService {
	return &commentServiceImpl{
		repository: repositories.NewCommentRepository(db),
	}
}

func (cc *commentServiceImpl) Create(request models.Comment) error {
	return cc.repository.Create(request)
}

func (cc *commentServiceImpl) List(userId any) ([]responses.CommentListResponse, error) {
	list, err := cc.repository.List(userId)
	if err != nil {
		return nil, err
	}

	var results []responses.CommentListResponse
	for _, v := range list {
		results = append(results, responses.CommentListResponse{
			ID:        int(v.ID),
			Message:   v.Message,
			PhotoID:   v.PhotoID,
			UserID:    v.UserID,
			CreatedAt: v.CreatedAt.Format(configs.TimeFormat),
			UpdatedAt: v.UpdatedAt.Format(configs.TimeFormat),
			User: responses.CommentListUserResponse{
				ID:       v.User.ID,
				Email:    v.User.Email,
				Username: v.User.Username,
			},
			Photo: responses.CommentListPhotoResponse{
				ID:       v.Photo.ID,
				Title:    v.Photo.Title,
				Caption:  v.Photo.Caption,
				PhotoUrl: v.Photo.PhotoUrl,
				UserId:   v.Photo.UserID,
			},
		})
	}

	return results, nil
}

func (cc *commentServiceImpl) Update(request requests.CommentUpdateRequest) (*models.Comment, error) {
	if isOwner := cc.repository.IsCommentOwner(request.UserID, request.ID); !isOwner {
		return nil, errors.New("unauthorized")
	}

	comment, err := cc.repository.Update(request)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (cc *commentServiceImpl) Delete(userId any, commentId string) error {
	if isOwner := cc.repository.IsCommentOwner(userId, commentId); !isOwner {
		return errors.New("unauthorized")
	}

	return cc.repository.Delete(commentId)
}
