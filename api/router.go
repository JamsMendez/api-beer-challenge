package api

import (
	"api-beer-challenge/internal/model"
	"api-beer-challenge/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

const prefixAPI = "/api"

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

	if beer == nil {
		return c.JSON(nil)
	}

	bJSON := BeerToJSON(beer)
	return c.JSON(bJSON)
}

func (r *routerHandler) getBeerBoxPrice(c *fiber.Ctx) error {
	const keyBeerID = "beerID"
	paramID, err := c.ParamsInt(keyBeerID)
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
	const quantityDefault = 6
	var quantity uint64

	queryQuantity := c.Query(keyQuantity)
	if queryQuantity == "" {
		quantity = quantityDefault
	} else {
		var value int
		value, err = strconv.Atoi(queryQuantity)
		if err != nil {
			msg := MessageJSON{Message: "invalid query 'quantity', must be a positive integer"}
			return c.Status(http.StatusBadRequest).JSON(msg)
		}

		quantity = uint64(value)
	}

	const keyCurrency = "currency"
	currency := c.Query(keyCurrency)
	if currency == "" {
		msg := MessageJSON{Message: "param required 'currency'"}
		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	beer, err := r.service.GetBeer(c.Context(), id)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	if beer == nil {
		return c.JSON(nil)
	}

	boxPrice, err := r.service.GetBeerBoxPrice(c.Context(), beer.ID, uint64(quantity), currency)
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
	beerInJSON := BeerInputJSON{}
	err := json.Unmarshal(c.Body(), &beerInJSON)
	if err != nil {
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

	bJSON := BeerToJSON(beer)
	return c.JSON(bJSON)
}
