package service

import (
	"context"

	"api-beer-challenge/internal/model"
	"api-beer-challenge/internal/repository"
	"api-beer-challenge/pkg/pagination"
)

//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	GetBeers(ctx context.Context, params *pagination.PaginationParams) (*pagination.Pagination, error)
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
