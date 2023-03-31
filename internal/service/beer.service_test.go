package service_test

import (
	"api-beer-challenge/internal/model"
	"api-beer-challenge/internal/service"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBeers(t *testing.T) {
	mockRepository.Test(t)

	ctx := context.Background()
	expected := []model.Beer{
		{
			ID:        1,
			Name:      "Corona",
			Brewery:   "Grupo Modelo",
			Country:   "Mexico",
			Price:     25.00,
			Currency:  "MXN",
			CreatedAt: createdAt,
			UpdatedAt: createdAt,
		},
		{
			ID:        2,
			Name:      "Estrella",
			Brewery:   "BeerHouse",
			Country:   "Mexico",
			Price:     20.00,
			Currency:  "MXN",
			CreatedAt: createdAt,
			UpdatedAt: createdAt,
		},
	}

	beers, err := serv.GetBeers(ctx)
	if err != nil {
		t.Fatalf("expected %v, got %v", expected, beers)
	}

	assert.Equal(t, expected, beers)
}

func TestGetBeer(t *testing.T) {
	mockRepository.Test(t)

	ctx := context.Background()
	beerID := uint64(1)

	expected := &model.Beer{
		ID:        1,
		Name:      "Corona",
		Brewery:   "Grupo Modelo",
		Country:   "Mexico",
		Price:     25.00,
		Currency:  "MXN",
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	beer, err := serv.GetBeer(ctx, beerID)
	if err != nil {
		t.Fatalf("expected %v, got %v", nil, err)
	}

	assert.Equal(t, expected, beer)
}

func TestGetBeerNotExists(t *testing.T) {
	mockRepository.Test(t)

	ctx := context.Background()
	beerID := uint64(99)

	beer, err := serv.GetBeer(ctx, beerID)
	if err != nil {
		t.Fatalf("expected %v, got %v", nil, err)
	}

	assert.Empty(t, beer)
}

func TestSaveBeer(t *testing.T) {
	mockRepository.Test(t)

	ctx := context.Background()

	expected := &model.Beer{
		ID:        1,
		Name:      "Corona",
		Brewery:   "Grupo Modelo",
		Country:   "Mexico",
		Price:     25.00,
		Currency:  "MXN",
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	beerInsert := model.InputBeer{
		Name:      "Corona",
		Brewery:   "Grupo Modelo",
		Country:   "Mexico",
		Price:     25.00,
		Currency:  "MXN",
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	beer, err := serv.SaveBeer(ctx, &beerInsert)
	if err != nil {
		t.Fatalf("expected %v, got %v", nil, err)
	}

	assert.Equal(t, expected, beer)
}

func TestGetBeerBoxPrice(t *testing.T) {
	mockRepository.Test(t)

	ctx := context.Background()

	expected := 300.00
	beerID := uint64(1)
	quantity := uint64(6)
	currency := "USD"

	price, err := serv.GetBeerBoxPrice(ctx, beerID, quantity, currency)
	if err != nil {
		t.Fatalf("expected %v, got %v", nil, err)
	}

	assert.Equal(t, expected, price)
}

func TestGetBeerBoxPriceQuantityDefault(t *testing.T) {
	ctx := context.Background()

	expected := 300.00
	beerID := uint64(1)
	quantity := uint64(0)
	currency := "USD"

	price, err := serv.GetBeerBoxPrice(ctx, beerID, quantity, currency)
	if err != nil {
		t.Fatalf("expected %v, got %v", nil, err)
	}

	assert.Equal(t, expected, price)
}

func TestGetBeerBoxPriceCurrencyEmpty(t *testing.T) {
	ctx := context.Background()

	expectedError := service.ErrCurrencyEmpty
	expected := 0.0

	beerID := uint64(1)
	quantity := uint64(0)
	currency := ""

	price, err := serv.GetBeerBoxPrice(ctx, beerID, quantity, currency)
	if err == nil {
		t.Fatalf("expected %v, got %v", expectedError, err)
	}

	assert.Equal(t, expectedError, err)
	assert.Equal(t, price, expected)
}
