package api

import "api-beer-challenge/internal/model"

const dateTimeFormat = "2006-01-02 15:04:05"

type beerJSON struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	Brewery   string  `json:"brewery"`
	Country   string  `json:"country"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type beerInputJSON struct {
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Country  string  `json:"country"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type beerBoxPriceJSON struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Currency string  `json:"currency"`
	Quantity uint64  `json:"quantity"`
	BoxPrice float64 `json:"box_price"`
}

type MessageJSON struct {
	Message string `json:"message"`
}

func beerToJSON(b *model.Beer) *beerJSON {
	bJSON := beerJSON{
		ID:        b.ID,
		Name:      b.Name,
		Brewery:   b.Brewery,
		Country:   b.Country,
		Price:     b.Price,
		Currency:  b.Currency,
		CreatedAt: b.CreatedAt.Format(dateTimeFormat),
		UpdatedAt: b.UpdatedAt.Format(dateTimeFormat),
	}

	return &bJSON
}
