package repository

import (
	"context"

	"github.com/Lupusdog/server-recruit-challenge-2024/model"
)

type AlbumRepository interface {
	GetAll(ctx context.Context) ([]*model.Album, error)
	Get(ctx context.Context, id model.AlbumID) (*model.Album, error)
	Add(ctx context.Context, album *model.Album) error
	Delete(ctx context.Context, id model.AlbumID) error
}