package service

import (
	"context"

	"api-beer-challenge/internal/entity"
	"api-beer-challenge/internal/model"
)

func (s *service) GetBeers(ctx context.Context) ([]model.Beer, error) {
	bb, err := s.repository.FindBeers(ctx)
	if err != nil {
		return nil, err
	}

	beers := []model.Beer{}

	for index := range bb {
		b := bb[index]
		beer := getBeerModel(&b)
		beers = append(beers, *beer)
	}

	return beers, nil
}

func (s *service) GetBeer(ctx context.Context, id uint64) (*model.Beer, error) {
	b, err := s.repository.FindBeerByID(ctx, id)
	if err != nil {
		return nil, err
	}

	beer := getBeerModel(b)
	return beer, nil
}

func (s *service) SaveBeer(ctx context.Context, input *model.BeerInput) (*model.Beer, error) {
	b, err := s.repository.InsertBeer(ctx, input)
	if err != nil {
		return nil, err
	}

	beer := getBeerModel(b)
	return beer, nil
}

func (s *service) GetBeerBoxPrice(ctx context.Context, id, quantity uint64, currency string) (float64, error) {
	return s.repository.FindBoxPriceBeer(ctx, id, quantity, currency)
}

func getBeerModel(b *entity.Beer) *model.Beer {
	return &model.Beer{
		ID:        b.ID,
		Name:      b.Name,
		Brewery:   b.Brewery,
		Country:   b.Country,
		Price:     b.Price,
		Currency:  b.Currency,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}
