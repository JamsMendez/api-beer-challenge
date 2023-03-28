package service

import (
	"context"

	"api-beer-challenge/internal/model"
	"api-beer-challenge/internal/repository"
)

type Service interface {
	GetBeers(ctx context.Context) ([]model.Beer, error)
	GetBeer(ctx context.Context, id uint64) (*model.Beer, error)
	SaveBeer(ctx context.Context, input *model.InputBeer) (*model.Beer, error)
	GetBeerBoxPrice(ctx context.Context, id, quantity uint64, currency string) (float64, error)
}

type service struct {
	repository repository.Repository
}

func New(repo repository.Repository) Service {
	return &service{
		repository: repo,
	}
}
