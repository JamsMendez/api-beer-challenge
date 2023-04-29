package service_test

import (
	"os"
	"testing"
	"time"

	"api-beer-challenge/internal/entity"
	"api-beer-challenge/internal/model"
	"api-beer-challenge/internal/service"

	mock "github.com/stretchr/testify/mock"
)

var mockRepository *MockRepository
var serv service.Service
var createdAt = time.Now().UTC()

func TestMain(m *testing.M) {
	mockRepository = &MockRepository{}

	// get all beers when are none
	// beersNone := []entity.Beer{}
	// mockRepository.On("FindBeers", mock.Anything).Return(beersNone, nil)

	beers := []entity.Beer{
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

	mockRepository.On("FindBeers", mock.Anything, mock.Anything, mock.Anything).Return(beers, nil)

	// get one beer
	beerID := uint64(1)
	beerIDNotExists := uint64(99)

	beer := &entity.Beer{
		ID:        1,
		Name:      "Corona",
		Brewery:   "Grupo Modelo",
		Country:   "Mexico",
		Price:     25.00,
		Currency:  "MXN",
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	mockRepository.On("FindBeerByID", mock.Anything, beerID).Return(beer, nil)
	mockRepository.On("FindBeerByID", mock.Anything, beerIDNotExists).Return(nil, nil)

	// save beer
	beerInsert := &model.InputBeer{
		Name:      "Corona",
		Brewery:   "Grupo Modelo",
		Country:   "Mexico",
		Price:     25.00,
		Currency:  "MXN",
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	mockRepository.On("InsertBeer", mock.Anything, beerInsert).Return(beer, nil)

	// get beer boxprice
	quantity := uint64(6)
	quantityTriggerDefault := uint64(0)
	currency := "USD"
	currencyEmpty := ""
	price := 300.00

	mockRepository.On("FindBoxPriceBeer", mock.Anything, beerID, quantity, currency).Return(price, nil)
	mockRepository.On("FindBoxPriceBeer", mock.Anything, beerID, quantityTriggerDefault, currency).Return(price, nil)
	mockRepository.On("FindBoxPriceBeer", mock.Anything, beerID, quantity, currencyEmpty).Return(nil, service.ErrCurrencyEmpty)

	serv = service.New(mockRepository)

	code := m.Run()
	os.Exit(code)
}
