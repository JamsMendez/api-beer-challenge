package api

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func validateReqGetBeer(c *fiber.Ctx) error {
	paramID, err := c.ParamsInt(keyParamID)
	if err != nil {
		msg := ErrorResponseJSON{
			Message: "invalid param 'id', must be a positive integer",
		}

		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	if paramID <= 0 {
		msg := ErrorResponseJSON{
			Message: "invalid param 'id', must be a positive integer",
		}

		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	id := uint64(paramID)
	c.Locals(keyParamID, id)

	return c.Next()
}

func validateReqGetBeerBoxPrice(c *fiber.Ctx) error {
	paramBeerID, err := c.ParamsInt(keyParamBeerID)
	if err != nil {
		msg := ErrorResponseJSON{
			Message: "invalid param 'id', must be a positive integer",
		}

		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	if paramBeerID <= 0 {
		msg := ErrorResponseJSON{
			Message: "invalid param 'id', must be a positive integer",
		}
		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	id := uint64(paramBeerID)

	const quantityDefault = 6
	var quantity uint64

	queryQuantity := c.Query(keyQueryQuantity)
	if queryQuantity == "" {
		quantity = quantityDefault
	} else {
		var value int
		value, err = strconv.Atoi(queryQuantity)
		if err != nil {
			msg := ErrorResponseJSON{
				Message: "invalid query 'quantity', must be a positive integer",
			}
			return c.Status(http.StatusBadRequest).JSON(msg)
		}

		quantity = uint64(value)
	}

	currency := c.Query(keyQueryCurrency)
	if currency == "" {
		msg := ErrorResponseJSON{
			Message: "param required 'currency'",
		}
		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	c.Locals(keyParamBeerID, id)
	c.Locals(keyQueryQuantity, quantity)
	c.Locals(keyQueryCurrency, currency)

	return c.Next()
}

func validateReqAddBeer(c *fiber.Ctx) error {
	var beerInJSON BeerNewJSON
	err := beerInJSON.New(c.Body())
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	err = beerInJSON.Validate()
	if err != nil {
		msg := ErrorResponseJSON{Message: err.Error()}
		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	c.Locals(keyInput, beerInJSON)
	return c.Next()
}
