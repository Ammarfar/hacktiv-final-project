package services

import (
	"finalproject/configs"
	"finalproject/models"
	"finalproject/repositories"
	"finalproject/responses"

	"gorm.io/gorm"
)

type PhotoService interface {
	Create(request models.Photo) error
	List(userId any) ([]responses.PhotoListResponse, error)
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
