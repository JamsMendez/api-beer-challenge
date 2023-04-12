package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"api-beer-challenge/api"
	"api-beer-challenge/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createdAt       = time.Unix(time.Now().Unix(), 0).UTC()
	createdAtFormat = createdAt.Format(api.DateTimeFormat)

	app *fiber.App
)

func TestMain(m *testing.M) {
	mockService := MockService{}

	routerHandler := api.NewRouterHandler(&mockService)
	app = fiber.New()
	api.SetUpRouters(app, routerHandler)

	// get all beers
	beers := []model.Beer{
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
	mockService.On("GetBeers", mock.Anything).Return(beers, nil)

	// get beer and get beer not found
	beerID := uint64(1)
	beerIDNotExists := uint64(99)

	beer := &model.Beer{
		ID:        1,
		Name:      "Corona",
		Brewery:   "Grupo Modelo",
		Country:   "Mexico",
		Price:     25.00,
		Currency:  "MXN",
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	mockService.On("GetBeer", mock.Anything, beerID).Return(beer, nil)
	mockService.On("GetBeer", mock.Anything, beerIDNotExists).Return(nil, nil)

	// get beer boxprice
	quantity := uint64(6)
	currency := "USD"
	price := 300.00

	mockService.On("GetBeerBoxPrice", mock.Anything, beerID, quantity, currency).Return(price, nil)

	// save beer
	mockService.On("SaveBeer", mock.Anything, mock.Anything).Return(beer, nil)

	code := m.Run()
	os.Exit(code)
}

func TestGetBeers(t *testing.T) {
	body := bytes.NewBuffer([]byte{})
	req := httptest.NewRequest("GET", "/api/beers", body)

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	expected := []api.BeerJSON{
		{
			ID:        1,
			Name:      "Corona",
			Brewery:   "Grupo Modelo",
			Country:   "Mexico",
			Price:     25.00,
			Currency:  "MXN",
			CreatedAt: createdAtFormat,
			UpdatedAt: createdAtFormat,
		},
		{
			ID:        2,
			Name:      "Estrella",
			Brewery:   "BeerHouse",
			Country:   "Mexico",
			Price:     20.00,
			Currency:  "MXN",
			CreatedAt: createdAtFormat,
			UpdatedAt: createdAtFormat,
		},
	}

	var beersJSON []api.BeerJSON
	err = json.NewDecoder(response.Body).Decode(&beersJSON)
	if err != nil {
		t.Fatalf("expected error nil, got %v", err)
	}

	current := len(beersJSON)
	size := len(expected)
	if current != size {
		t.Fatalf("expected beers size %v, got %v", size, current)
	}

	assert.Equal(t, expected, beersJSON)
}

func TestGetBeer(t *testing.T) {
	body := bytes.NewBuffer([]byte{})

	ID := 1
	target := fmt.Sprintf("/api/beers/%d", ID)
	req := httptest.NewRequest("GET", target, body)

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	expected := api.BeerJSON{
		ID:        1,
		Name:      "Corona",
		Brewery:   "Grupo Modelo",
		Country:   "Mexico",
		Price:     25.00,
		Currency:  "MXN",
		CreatedAt: createdAtFormat,
		UpdatedAt: createdAtFormat,
	}

	var beerJSON api.BeerJSON
	err = json.NewDecoder(response.Body).Decode(&beerJSON)
	if err != nil {
		t.Fatalf("expected error nil, got %v", err)
	}

	assert.Equal(t, expected, beerJSON)
}

func TestGetBeerNotFound(t *testing.T) {
	body := bytes.NewBuffer([]byte{})

	ID := 99
	target := fmt.Sprintf("/api/beers/%d", ID)
	req := httptest.NewRequest("GET", target, body)

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var beerJSON api.BeerJSON
	err = json.NewDecoder(response.Body).Decode(&beerJSON)
	if err != nil {
		t.Fatalf("expected error nil, got %v", err)
	}

	if !assert.Empty(t, beerJSON) {
		t.Fatalf("expected JSON %v, got %v", nil, beerJSON)
	}
}

func TestGetBeerBoxPrice(t *testing.T) {
	body := bytes.NewBuffer([]byte{})

	ID := 1
	quality := 6
	currency := "USD"

	target := fmt.Sprintf("/api/beers/%d/boxprice?quantity=%d&currency=%s", ID, quality, currency)
	r, err := url.Parse(target)
	if err != nil {
		t.Fatalf("expected url parse error nil, got %v", err)
	}

	req := httptest.NewRequest("GET", r.String(), body)

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	expected := api.BeerBoxPriceJSON{
		ID:       1,
		Name:     "Corona",
		Brewery:  "Grupo Modelo",
		Currency: "USD",
		Quantity: 6,
		BoxPrice: 300,
	}

	beerBoxPriceJSON := api.BeerBoxPriceJSON{}
	err = json.NewDecoder(response.Body).Decode(&beerBoxPriceJSON)
	if err != nil {
		t.Fatalf("expected error nil, got %v", err)
	}

	assert.Equal(t, expected, beerBoxPriceJSON)
}

func TestGetBeerParamIDFailed(t *testing.T) {
	body := bytes.NewBuffer([]byte{})

	ID := "XX"
	target := fmt.Sprintf("/api/beers/%s", ID)
	req := httptest.NewRequest("GET", target, body)

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestGetBeerBoxPriceParamIDNegative(t *testing.T) {
	body := bytes.NewBuffer([]byte{})

	ID := -1
	quality := 6
	currency := "USD"

	target := fmt.Sprintf("/api/beers/%d/boxprice?quantity=%d&currency=%s", ID, quality, currency)
	req := httptest.NewRequest("GET", target, body)

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestGetBeerBoxPriceParamIDFaild(t *testing.T) {
	body := bytes.NewBuffer([]byte{})

	ID := "XXX"
	quality := 6
	currency := "USD"

	target := fmt.Sprintf("/api/beers/%s/boxprice?quantity=%d&currency=%s", ID, quality, currency)
	req := httptest.NewRequest("GET", target, body)

	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestGetBeerBoxPriceParamQuantityFailed(t *testing.T) {
	body := bytes.NewBuffer([]byte{})

	ID := 1
	quantity := "XX"
	currency := "USD"

	u := fmt.Sprintf("/api/beers/%d/boxprice?quantity=%s&currency=%s", ID, quantity, currency)
	target, err := url.Parse(u)
	if err != nil {
		t.Fatalf("expected url parse error nil, got %v", err)
	}

	req := httptest.NewRequest("GET", target.String(), body)
	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestGetBeerBoxPriceParamCurrencyFailed(t *testing.T) {
	body := bytes.NewBuffer([]byte{})

	ID := 1
	quantity := 6
	currency := ""

	u := fmt.Sprintf("/api/beers/%d/boxprice?quantity=%d&currency=%s", ID, quantity, currency)
	target, err := url.Parse(u)
	if err != nil {
		t.Fatalf("expected url parse error nil, got %v", err)
	}

	req := httptest.NewRequest("GET", target.String(), body)
	response, err := app.Test(req)
	if err != nil {
		t.Fatalf("expected app test error nil, got %v", err)
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestAddBeer(t *testing.T) {
	input := &api.BeerNewJSON{
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

	assert.Equal(t, http.StatusOK, response.StatusCode)

	expected := api.BeerJSON{
		ID:        1,
		Name:      "Corona",
		Brewery:   "Grupo Modelo",
		Country:   "Mexico",
		Price:     25.00,
		Currency:  "MXN",
		CreatedAt: createdAtFormat,
		UpdatedAt: createdAtFormat,
	}

	var beerJSON api.BeerJSON
	err = json.NewDecoder(response.Body).Decode(&beerJSON)
	if err != nil {
		t.Fatalf("expected error nil, got %v", err)
	}

	assert.Equal(t, expected, beerJSON)
}
