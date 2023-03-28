package model

import "time"

type Beer struct {
	ID        uint64
	Name      string
	Brewery   string
	Country   string
	Price     float64
	Currency  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InputBeer struct {
	Name      string
	Brewery   string
	Country   string
	Price     float64
	Currency  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InputUBeer struct {
	Name      InputU
	Brewery   InputU
	Country   InputU
	Price     InputU
	Currency  InputU
	CreatedAt InputU
	UpdatedAt InputU
}
