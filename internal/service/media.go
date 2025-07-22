package service

import (
	"context"
	"database/sql"
	"path"
	"time"

	"github.com/LouisFernando1204/golang-restapi.git/domain"
	"github.com/LouisFernando1204/golang-restapi.git/dto"
	"github.com/LouisFernando1204/golang-restapi.git/internal/config"
	"github.com/google/uuid"
)

type mediaService struct {
	conf            *config.Config
	mediaRepository domain.MediaRepository
}

func NewMedia(conf *config.Config, mediaRepository domain.MediaRepository) domain.MediaService {
	return &mediaService{
		conf:            conf,
		mediaRepository: mediaRepository,
	}
}

// Create implements domain.MediaService.
func (m *mediaService) Create(ctx context.Context, req dto.CreateMediaRequest) (dto.MediaData, error) {
	media := domain.Media{
		Id:        uuid.NewString(),
		Path:      req.Path,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	err := m.mediaRepository.Save(ctx, &media)
	if err != nil {
		return dto.MediaData{}, err
	}

	url := path.Join(m.conf.Server.Asset, media.Path)
	return dto.MediaData{
		Id:   media.Id,
		Path: media.Path,
		Url:  url,
	}, nil
}
