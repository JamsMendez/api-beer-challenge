package repository

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"api-beer-challenge/internal/entity"
	"api-beer-challenge/internal/model"
)

const (
	queryInsertBeer = `
	INSERT INTO
		beers
		(
			name,
			brewery,
			country,
			price,
			currency,
			created_at,
			updated_at
		)
	VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)
	RETURNING id;
	`

	queryFindBeers = `
		SELECT
			id,
			name,
			brewery,
			country,
			price,
			currency,
			created_at,
			updated_at
		FROM
			beers;
	`
	queryFindBeerByID = `
		SELECT
			id,
			name,
			brewery,
			country,
			price,
			currency,
			created_at,
			updated_at
		FROM
			beers
		WHERE
			id = $1;
	`

	queryUpdateBeerByID = `
		UPDATE
			beers
		SET
			%s
		WHERE id = $1;
	`

	queryDeleteBeerByID = `
		DELETE
		FROM
			beers
		WHERE
			id = $1;
	`
)

const keyAPILayer = "..."

var ErrRequestInvalid = errors.New("request invalid for api")

type ResponseJSON struct {
	Info struct {
		Quote     float64 `json:"quote"`
		Timestamp int     `json:"timestamp"`
	} `json:"info"`
	Query struct {
		Amount int    `json:"amount"`
		From   string `json:"from"`
		To     string `json:"to"`
	} `json:"query"`
	Result  float64 `json:"result"`
	Success bool    `json:"success"`
}

func (r *repository) FindBeers(ctx context.Context) ([]entity.Beer, error) {
	beers := []entity.Beer{}

	err := r.db.SelectContext(ctx, &beers, queryFindBeers)
	if err != nil {
		return nil, err
	}

	return beers, err
}

func (r *repository) FindBeerByID(ctx context.Context, id uint64) (*entity.Beer, error) {
	beer := entity.Beer{}

	err := r.db.GetContext(ctx, &beer, queryFindBeerByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &beer, err
}

func (r *repository) FindBoxPriceBeer(ctx context.Context, id, quantity uint64, to string) (float64, error) {
	var price float64
	b, err := r.FindBeerByID(ctx, id)
	if err != nil {
		return price, err
	}

	if b == nil {
		return price, ErrNotFoundEntity
	}

	if quantity == 0 {
		quantity = 6
	}

	amount := b.Price * float64(quantity)
	price, err = getConvertCurrent(b.Currency, to, amount)
	if err != nil {
		return price, err
	}

	return price, nil
}

func (r *repository) FindBoxPriceBeerFake(ctx context.Context, id, quantity uint64, to string) (float64, error) {
	var price float64
	b, err := r.FindBeerByID(ctx, id)
	if err != nil {
		return price, err
	}

	if quantity == 0 {
		quantity = 6
	}

	amount := b.Price * float64(quantity)
	price, err = getConvertCurrentFake(b.Currency, to, amount)
	if err != nil {
		return price, err
	}

	return price, nil
}

func (r *repository) InsertBeer(ctx context.Context, input *model.InputBeer) (*entity.Beer, error) {
	beerID, err := r.insertEntity(
		ctx,
		queryInsertBeer,
		input.Name,
		input.Brewery,
		input.Country,
		input.Price,
		input.Currency,
		input.CreatedAt,
		input.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	beer, err := r.FindBeerByID(ctx, beerID)
	if err != nil {
		return nil, err
	}

	return beer, nil
}

func (r *repository) UpdateBeerByID(ctx context.Context, id uint64, input *model.InputUBeer) (*entity.Beer, error) {
	var fields []any
	var columns []string
	// first argument is ID
	var numArg = 2

	if input.Name.Valid {
		fields = append(fields, input.Name.Value)
		columns = append(
			columns,
			fmt.Sprintf("name = $%d", numArg),
		)
		numArg++
	}

	if input.Brewery.Valid {
		fields = append(fields, input.Brewery.Value)
		columns = append(
			columns,
			fmt.Sprintf("brewery = $%d", numArg),
		)
		numArg++
	}

	if input.Country.Valid {
		fields = append(fields, input.Country.Value)
		columns = append(
			columns,
			fmt.Sprintf("country = $%d", numArg),
		)
		numArg++
	}

	if input.Price.Valid {
		fields = append(fields, input.Price.Value)
		columns = append(
			columns,
			fmt.Sprintf("price = $%d", numArg),
		)
		numArg++
	}

	if input.Currency.Valid {
		fields = append(fields, input.Currency.Value)
		columns = append(
			columns,
			fmt.Sprintf("currency = $%d", numArg),
		)
		numArg++
	}

	if input.CreatedAt.Valid {
		fields = append(fields, input.CreatedAt.Value)
		columns = append(
			columns,
			fmt.Sprintf("created_at = $%d", numArg),
		)
		numArg++
	}

	if input.UpdatedAt.Valid {
		fields = append(fields, input.UpdatedAt.Value)
		columns = append(
			columns,
			fmt.Sprintf("updated_at = $%d", numArg),
		)
	}

	values := strings.Join(columns, ", ")
	query := fmt.Sprintf(queryUpdateBeerByID, values)

	err := r.updateEntity(ctx, query, id, fields...)
	if err != nil {
		return nil, err
	}

	client, err := r.FindBeerByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return client, err
}

func (r *repository) DeleteBeerByID(ctx context.Context, id uint64) error {
	err := r.deleteEntity(ctx, queryDeleteBeerByID, id)
	return err
}

func (r *repository) RestartTable(ctx context.Context, src string) error {
	buffer, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	queryRestartTableBeers := string(buffer)
	_, err = r.db.ExecContext(ctx, queryRestartTableBeers)
	return err
}

func getConvertCurrentFake(from, to string, amount float64) (float64, error) {
	var err error

	if from == to {
		return amount, err
	}

	return amount * 2, err
}

func getConvertCurrent(from, to string, amount float64) (float64, error) {
	url := fmt.Sprintf(
		"https://api.apilayer.com/currency_data/convert?from=%s&to=%s&amount=%.2f",
		from,
		to,
		amount,
	)

	const seconds = 5

	var result float64
	timeout := seconds * time.Second

	body := bytes.NewBuffer([]byte{})
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, body)
	if err != nil {
		return result, err
	}

	req.Header.Set("apiKey", keyAPILayer)

	res, err := client.Do(req)
	if err != nil {
		return result, err
	}

	if res.StatusCode != http.StatusOK {
		return result, ErrRequestInvalid
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	buffer, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	var resJSON ResponseJSON
	err = json.Unmarshal(buffer, &resJSON)
	if err != nil {
		return result, err
	}

	result = resJSON.Result

	return result, err
}
