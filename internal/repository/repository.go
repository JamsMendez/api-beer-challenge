package repository

import (
	"context"

	"api-beer-challenge/internal/entity"
	"api-beer-challenge/internal/model"

	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	Count(ctx context.Context) (uint32, error)
	FindBeers(ctx context.Context, skip, limit uint32) ([]entity.Beer, error)
	FindBeerByID(ctx context.Context, id uint64) (*entity.Beer, error)
	FindBoxPriceBeer(
		ctx context.Context,
		id,
		quantity uint64,
		currency string,
	) (float64, error)

	InsertBeer(ctx context.Context, input *model.InputBeer) (*entity.Beer, error)
	UpdateBeerByID(ctx context.Context, id uint64, input *model.InputUBeer) (*entity.Beer, error)
	DeleteBeerByID(ctx context.Context, id uint64) error

	RestartTable(ctx context.Context, src string) error
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repository{db}
}
