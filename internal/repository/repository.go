package repository

import (
	"context"

	"api-beer-challenge/internal/entity"
	"api-beer-challenge/internal/model"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindBeers(ctx context.Context) ([]entity.Beer, error)
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

func New(db *sqlx.DB) *repository {
	return &repository{db}
}
