package api

import (
	"api-beer-challenge/database"
	"api-beer-challenge/internal/repository"
	"api-beer-challenge/internal/service"
	"api-beer-challenge/settings"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestAPI(t *testing.T) {
	s, err := settings.New()
	if err != nil {
		t.Fatalf("expected settings error nil, got %v", err)
	}

	ctx := context.Background()
	db, err := database.GetConnection(ctx, s)
	if err != nil {
		t.Fatalf("expected database error nil, got %v", err)
	}

	repo := repository.New(db)
	err = repo.RestartTable(ctx, "./../database/schema.sql")
	if err != nil {
		t.Fatalf("expected restart database error nil, got %v", err)
	}

	serv := service.New(repo)

	routerHandler := newRouterHandler(serv)
	app := fiber.New()
	setUpRouters(app, routerHandler)

	t.Run("get list beers", func(t *testing.T) {
		getBeers(app, t)
	})

	t.Run("add beer", func(t *testing.T) {
		addBeer(app, t)
	})

	t.Run("get beer", func(t *testing.T) {
		getBeer(app, t)
	})

	t.Run("get beer boxprice", func(t *testing.T) {
		getBeerBoxPrice(app, t)
	})

}

func getBeers(app *fiber.App, t *testing.T) {
	body := bytes.NewBuffer([]byte{})
	req := httptest.NewRequest("GET", "/api/beers", body)

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	beersJSON := []beerJSON{}
	err = json.NewDecoder(response.Body).Decode(&beersJSON)
	if err != nil {
		t.Fatalf("expected error nil, got %v", err)
	}

	size := len(beersJSON)
	if size != 0 {
		t.Fatalf("expected beers size 0, got %v", size)
	}
}

func getBeer(app *fiber.App, t *testing.T) {
	body := bytes.NewBuffer([]byte{})

	ID := 1
	target := fmt.Sprintf("/api/beers/%d", ID)
	req := httptest.NewRequest("GET", target, body)

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	bJSON := beerJSON{}
	err = json.NewDecoder(response.Body).Decode(&bJSON)
	if err != nil {
		t.Fatalf("expected error nil, got %v", err)
	}

	bOutJSON := beerJSON{
		ID:       1,
		Name:     "Corona",
		Brewery:  "Grupo Modelo",
		Country:  "Mexico",
		Price:    25.00,
		Currency: "MXN",
	}

	assertBeer(&bOutJSON, &bJSON, t)
}

func getBeerBoxPrice(app *fiber.App, t *testing.T) {
	body := bytes.NewBuffer([]byte{})

	ID := 1
	quality := 6
	currency := "USD"

	target := fmt.Sprintf("/api/beers/%d/boxprice?quality=%d&currency=%s", ID, quality, currency)
	req := httptest.NewRequest("GET", target, body)

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	bBoxPriceJSON := beerBoxPriceJSON{}
	err = json.NewDecoder(response.Body).Decode(&bBoxPriceJSON)
	if err != nil {
		t.Fatalf("expected error nil, got %v", err)
	}

	bOutJSON := beerBoxPriceJSON{
		ID:       1,
		Name:     "Corona",
		Brewery:  "Grupo Modelo",
		Currency: "USD",
		Quantity: 6,
		BoxPrice: 300,
	}

	assertBeerBoxPrice(&bOutJSON, &bBoxPriceJSON, t)
}

func addBeer(app *fiber.App, t *testing.T) {
	input := beerInputJSON{
		Name:     "Corona",
		Brewery:  "Grupo Modelo",
		Country:  "Mexico",
		Price:    25.00,
		Currency: "MXN",
	}

	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(&input)
	if err != nil {
		t.Fatalf("expected json encode error nil, got %v", err)
	}

	req := httptest.NewRequest("POST", "/api/beers", body)

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	bJSON := beerJSON{}
	err = json.NewDecoder(response.Body).Decode(&bJSON)
	if err != nil {
		t.Fatalf("expected error nil, got %v", err)
	}

	bOutJSON := beerJSON{
		ID:       1,
		Name:     "Corona",
		Brewery:  "Grupo Modelo",
		Country:  "Mexico",
		Price:    25.00,
		Currency: "MXN",
	}

	assertBeer(&bOutJSON, &bJSON, t)
}

func assertBeerBoxPrice(want, got *beerBoxPriceJSON, t testing.TB) {
	if want.ID != got.ID {
		t.Fatalf("expected ID %d, got ID %d", want.ID, got.ID)
	}

	if want.Name != got.Name {
		t.Fatalf("expected Name %s, got Name %s", want.Name, got.Name)
	}

	if want.Brewery != got.Brewery {
		t.Fatalf("expected Brewery %s, got Brewery %s", want.Brewery, got.Brewery)
	}

	if want.Currency != got.Currency {
		t.Fatalf("expected Currency %s, got Currency %s", want.Currency, got.Currency)
	}

	if want.BoxPrice != got.BoxPrice {
		t.Fatalf("expected BoxPrice %.2f, got BoxPrice %.2f", want.BoxPrice, got.BoxPrice)
	}

	if want.Quantity != got.Quantity {
		t.Fatalf("expected Quantity %d, got Quantity %d", want.Quantity, got.Quantity)
	}


	//TODO: how?
	/*
		if want.CreatedAt != got.CreatedAt {
			t.Fatalf("expected CreatedAt %s, got CreatedAt %s", want.CreatedAt, got.CreatedAt)
		}

		if want.UpdatedAt != got.UpdatedAt {
			t.Fatalf("expected UpdatedAt %s, got UpdatedAt %s", want.UpdatedAt, got.UpdatedAt)
		}
	*/
}

func assertBeer(want, got *beerJSON, t testing.TB) {
	if want.ID != got.ID {
		t.Fatalf("expected ID %d, got ID %d", want.ID, got.ID)
	}

	if want.Name != got.Name {
		t.Fatalf("expected Name %s, got Name %s", want.Name, got.Name)
	}

	if want.Brewery != got.Brewery {
		t.Fatalf("expected Brewery %s, got Brewery %s", want.Brewery, got.Brewery)
	}

	if want.Country != got.Country {
		t.Fatalf("expected Country %s, got Country %s", want.Country, got.Country)
	}

	if want.Price != got.Price {
		t.Fatalf("expected Price %.2f, got Price %.2f", want.Price, got.Price)
	}

	if want.Currency != got.Currency {
		t.Fatalf("expected Currency %s, got Currency %s", want.Currency, got.Currency)
	}

	//TODO: how?
	/*
		if want.CreatedAt != got.CreatedAt {
			t.Fatalf("expected CreatedAt %s, got CreatedAt %s", want.CreatedAt, got.CreatedAt)
		}

		if want.UpdatedAt != got.UpdatedAt {
			t.Fatalf("expected UpdatedAt %s, got UpdatedAt %s", want.UpdatedAt, got.UpdatedAt)
		}
	*/
}
