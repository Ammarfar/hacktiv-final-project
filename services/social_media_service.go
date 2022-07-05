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

type SocialMediaService interface {
	Create(request models.SocialMedia) error
	List(userId any) ([]responses.SocialMediaListResponse, error)
	Update(request requests.SocialMediaUpdateRequest) (*models.SocialMedia, error)
	Delete(userId any, socialMediaId string) error
}

type socialMediaServiceImpl struct {
	repository repositories.SocialMediaRepository
}

func NewSocialMediaService(db *gorm.DB) SocialMediaService {
	return &socialMediaServiceImpl{
		repository: repositories.NewSocialMediaRepository(db),
	}
}

func (pc *socialMediaServiceImpl) Create(request models.SocialMedia) error {
	return pc.repository.Create(request)
}

func (pc *socialMediaServiceImpl) List(userId any) ([]responses.SocialMediaListResponse, error) {
	list, err := pc.repository.List(userId)
	if err != nil {
		return nil, err
	}

	var results []responses.SocialMediaListResponse
	for _, v := range list {
		results = append(results, responses.SocialMediaListResponse{
			ID:             int(v.ID),
			Name:           v.Name,
			SocialMediaUrl: v.SocialMediaUrl,
			UserID:         int(v.UserID),
			CreatedAt:      v.CreatedAt.Format(configs.TimeFormat),
			UpdatedAt:      v.UpdatedAt.Format(configs.TimeFormat),
			User: responses.SocialMediaListUserResponse{
				ID:       v.User.ID,
				Username: v.User.Username,
			},
		})
	}

	return results, nil
}

func (pc *socialMediaServiceImpl) Update(request requests.SocialMediaUpdateRequest) (*models.SocialMedia, error) {
	if isOwner := pc.repository.IsSocialMediaOwner(request.UserID, request.ID); !isOwner {
		return nil, errors.New("unauthorized")
	}

	socialMedia, err := pc.repository.Update(request)
	if err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (pc *socialMediaServiceImpl) Delete(userId any, socialMediaId string) error {
	if isOwner := pc.repository.IsSocialMediaOwner(userId, socialMediaId); !isOwner {
		return errors.New("unauthorized")
	}

	return pc.repository.Delete(socialMediaId)
}
