package api

import (
	"net/http"
	"time"

	"api-beer-challenge/internal/model"
	"api-beer-challenge/internal/service"

	"github.com/gofiber/fiber/v2"
)

type routerHandler struct {
	service service.Service
}

func NewRouterHandler(s service.Service) *routerHandler {
	return &routerHandler{s}
}

func (r *routerHandler) getBeers(c *fiber.Ctx) error {
	bb, err := r.service.GetBeers(c.Context())
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	beers := []BeerJSON{}

	for index := range bb {
		b := bb[index]
		beer := BeerToJSON(&b)
		beers = append(beers, *beer)
	}

	return c.JSON(beers)
}

func (r *routerHandler) getBeer(c *fiber.Ctx) error {
	id, ok := c.Locals(keyParamID).(uint64)
	if !ok {
		return c.SendStatus(http.StatusBadRequest)
	}

	beer, err := r.service.GetBeer(c.Context(), id)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	if beer == nil {
		return c.JSON(nil)
	}

	bJSON := BeerToJSON(beer)
	return c.JSON(bJSON)
}

func (r *routerHandler) getBeerBoxPrice(c *fiber.Ctx) error {
	id, ok := c.Locals(keyParamBeerID).(uint64)
	if !ok {
		return c.SendStatus(http.StatusBadRequest)
	}

	quantity, ok := c.Locals(keyQueryQuantity).(uint64)
	if !ok {
		return c.SendStatus(http.StatusBadRequest)
	}

	currency, ok := c.Locals(keyQueryCurrency).(string)
	if !ok {
		return c.SendStatus(http.StatusBadRequest)
	}

	beer, err := r.service.GetBeer(c.Context(), id)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	if beer == nil {
		return c.JSON(nil)
	}

	boxPrice, err := r.service.GetBeerBoxPrice(c.Context(), beer.ID, quantity, currency)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	beerJSON := BeerBoxPriceJSON{
		ID:       beer.ID,
		Name:     beer.Name,
		Brewery:  beer.Brewery,
		Currency: currency,
		Quantity: quantity,
		BoxPrice: boxPrice,
	}

	return c.JSON(beerJSON)
}

func (r *routerHandler) addBeer(c *fiber.Ctx) error {
	beerInJSON, ok := c.Locals(keyInput).(BeerNewJSON)
	if !ok {
		return c.SendStatus(http.StatusBadRequest)
	}

	createdAt := time.Unix(time.Now().Unix(), 0).UTC()

	beerInput := model.InputBeer{
		Name:      beerInJSON.Name,
		Brewery:   beerInJSON.Brewery,
		Country:   beerInJSON.Country,
		Price:     beerInJSON.Price,
		Currency:  beerInJSON.Currency,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	beer, err := r.service.SaveBeer(c.Context(), &beerInput)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	beerJSON := BeerToJSON(beer)
	return c.JSON(&beerJSON)
}
