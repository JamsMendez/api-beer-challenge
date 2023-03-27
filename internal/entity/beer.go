package entity

import "time"

type Beer struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	Brewery   string    `db:"brewery"`
	Country   string    `db:"country"`
	Price     float64   `db:"price"`
	Currency  string    `db:"currency"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
