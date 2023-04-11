package api

import (
	"encoding/json"
	"fmt"

	"api-beer-challenge/internal/model"

	"github.com/go-playground/validator/v10"
)

const DateTimeFormat = "2006-01-02 15:04:05"

const (
	keyParamID       = "id"
	keyParamBeerID   = "beerID"
	keyQueryQuantity = "quantity"
	keyQueryCurrency = "currency"

	keyInput = "input"
)

type BeerJSON struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	Brewery   string  `json:"brewery"`
	Country   string  `json:"country"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type BeerNewJSON struct {
	Name     string  `json:"name" validate:"required,min=2"`
	Brewery  string  `json:"brewery" validate:"required,min=2"`
	Country  string  `json:"country" validate:"required,min=2"`
	Price    float64 `json:"price" validate:"required,number,gte=0"`
	Currency string  `json:"currency" validate:"required,len=3"`

	validator *validator.Validate
}

func (b *BeerNewJSON) New(buffer []byte) error {
	err := json.Unmarshal(buffer, b)
	if err != nil {
		return err
	}

	b.validator = validator.New()
	return nil
}

func (b *BeerNewJSON) Validate() error {
	err := b.validator.Struct(b)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return fmt.Errorf("invalid beer json")
		}

		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("field %s invalid", err.Field())
		}
	}

	return err
}

type BeerBoxPriceJSON struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Currency string  `json:"currency"`
	Quantity uint64  `json:"quantity"`
	BoxPrice float64 `json:"box_price"`
}

type ErrorResponseJSON struct {
	Message string `json:"message"`
}

func BeerToJSON(b *model.Beer) *BeerJSON {
	return &BeerJSON{
		ID:        b.ID,
		Name:      b.Name,
		Brewery:   b.Brewery,
		Country:   b.Country,
		Price:     b.Price,
		Currency:  b.Currency,
		CreatedAt: b.CreatedAt.Format(DateTimeFormat),
		UpdatedAt: b.UpdatedAt.Format(DateTimeFormat),
	}
}
