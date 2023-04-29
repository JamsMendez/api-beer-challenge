package service

import (
	"context"
	"errors"
	"fmt"

	"api-beer-challenge/internal/entity"
	"api-beer-challenge/internal/model"
	"api-beer-challenge/pkg/pagination"
)

var (
	ErrCurrencyEmpty = errors.New("currency value is empty")
)

func (s *service) GetBeers(ctx context.Context, params *pagination.PaginationParams) (*pagination.Pagination, error) {
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, err
	}

	var beers []model.Beer

	if count == 0 {
		beerPagination := pagination.New(params.Page, params.PerPage, count)
		beerPagination.Items = beers
		return beerPagination, nil
	}

	skip := params.PerPage * (params.Page - 1)
	limit := params.PerPage

	fmt.Println(skip, limit)

	bb, err := s.repository.FindBeers(ctx, skip, limit)
	if err != nil {
		return nil, err
	}

	for index := range bb {
		b := bb[index]
		beer := getBeerModel(&b)
		beers = append(beers, *beer)
	}

	beerPagination := pagination.New(params.Page, params.PerPage, count)
	beerPagination.Size = uint32(len(beers))
	beerPagination.Items = beers

	return beerPagination, nil
}

func (s *service) GetBeer(ctx context.Context, id uint64) (*model.Beer, error) {
	b, err := s.repository.FindBeerByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if b == nil {
		return nil, nil
	}

	beer := getBeerModel(b)
	return beer, nil
}

func (s *service) SaveBeer(ctx context.Context, input *model.InputBeer) (*model.Beer, error) {
	b, err := s.repository.InsertBeer(ctx, input)
	if err != nil {
		return nil, err
	}

	beer := getBeerModel(b)
	return beer, nil
}

func (s *service) GetBeerBoxPrice(ctx context.Context, id, quantity uint64, currency string) (float64, error) {
	if currency == "" {
		return 0, ErrCurrencyEmpty
	}

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
