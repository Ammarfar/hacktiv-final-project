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

type PhotoService interface {
	Create(request models.Photo) error
	List(userId any) ([]responses.PhotoListResponse, error)
	Update(request requests.PhotoUpdateRequest) (*models.Photo, error)
	Delete(userId any, photoId string) error
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

func (pc *photoServiceImpl) List(userId any) ([]responses.PhotoListResponse, error) {
	list, err := pc.repository.List(userId)
	if err != nil {
		return nil, err
	}

	var results []responses.PhotoListResponse
	for _, v := range list {
		results = append(results, responses.PhotoListResponse{
			ID:        int(v.ID),
			Title:     v.Title,
			Caption:   v.Caption,
			PhotoUrl:  v.PhotoUrl,
			UserID:    int(v.UserID),
			CreatedAt: v.CreatedAt.Format(configs.TimeFormat),
			UpdatedAt: v.UpdatedAt.Format(configs.TimeFormat),
			User: responses.PhotoListUserResponse{
				Email:    v.User.Email,
				Username: v.User.Username,
			},
		})
	}

	return results, nil
}

func (pc *photoServiceImpl) Update(request requests.PhotoUpdateRequest) (*models.Photo, error) {
	if isOwner := pc.repository.IsPhotoOwner(request.UserID, request.ID); !isOwner {
		return nil, errors.New("unauthorized")
	}

	photo, err := pc.repository.Update(request)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (pc *photoServiceImpl) Delete(userId any, photoId string) error {
	if isOwner := pc.repository.IsPhotoOwner(userId, photoId); !isOwner {
		return errors.New("unauthorized")
	}

	return pc.repository.Delete(photoId)
}
