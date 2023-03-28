package api

import (
	"api-beer-challenge/internal/model"
	"api-beer-challenge/internal/service"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

const prefixAPI = "/api"

type routerHandler struct {
	service service.Service
}

func newRouterHandler(s service.Service) *routerHandler {
	return &routerHandler{s}
}

func (r *routerHandler) getBeers(c *fiber.Ctx) error {
	bb, err := r.service.GetBeers(c.Context())
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	beers := []beerJSON{}

	for index := range bb {
		b := bb[index]
		beer := beerToJSON(&b)
		beers = append(beers, *beer)
	}

	return c.JSON(beers)
}

func (r *routerHandler) getBeer(c *fiber.Ctx) error {
	paramID, err := c.ParamsInt("id")
	if err != nil {
		msg := MessageJSON{Message: "invalid param 'id', must be a positive integer"}
		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	if paramID <= 0 {
		msg := MessageJSON{Message: "invalid param 'id', must be a positive integer"}
		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	id := uint64(paramID)
	beer, err := r.service.GetBeer(c.Context(), id)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	bJSON := beerToJSON(beer)
	return c.JSON(bJSON)
}

func (r *routerHandler) getBeerBoxPrice(c *fiber.Ctx) error {
	paramID, err := c.ParamsInt("beerID")
	if err != nil {
		msg := MessageJSON{Message: "invalid param 'id', must be a positive integer"}
		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	if paramID <= 0 {
		msg := MessageJSON{Message: "invalid param 'id', must be a positive integer"}
		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	id := uint64(paramID)

	const keyQuantity = "quantity"
	const keyCurrency = "currency"
	const defaultCurrency = "USD"
	const defaultQuantity = 6

	quantity := c.QueryInt(keyQuantity, defaultQuantity)
	currency := c.Query(keyCurrency, defaultCurrency)

	beer, err := r.service.GetBeer(c.Context(), id)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	boxPrice, err := r.service.GetBeerBoxPrice(c.Context(), beer.ID, uint64(quantity), currency)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	beerJSON := beerBoxPriceJSON{
		ID:       beer.ID,
		Name:     beer.Name,
		Brewery:  beer.Brewery,
		Currency: currency,
		Quantity: uint64(quantity),
		BoxPrice: boxPrice,
	}

	return c.JSON(beerJSON)
}

func (r *routerHandler) addBeer(c *fiber.Ctx) error {
	beerInJSON := beerInputJSON{}
	err := json.Unmarshal(c.Body(), &beerInJSON)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	now := time.Now().UTC()
	beerInput := model.InputBeer{
		Name:      beerInJSON.Name,
		Brewery:   beerInJSON.Brewery,
		Country:   beerInJSON.Country,
		Price:     beerInJSON.Price,
		Currency:  beerInJSON.Currency,
		CreatedAt: now,
		UpdatedAt: now,
	}

	beer, err := r.service.SaveBeer(c.Context(), &beerInput)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	bJSON := beerToJSON(beer)
	return c.JSON(bJSON)
}
