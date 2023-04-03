package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
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
	c.Locals(keyParamBeerID, id)

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

	c.Locals(keyQueryQuantity, quantity)

	currency := c.Query(keyQueryCurrency)
	if currency == "" {
		msg := ErrorResponseJSON{
			Message: "param required 'currency'",
		}
		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	c.Locals(keyQueryCurrency, currency)

	return c.Next()
}

func validateReqAddBeer(c *fiber.Ctx) error {
	var beerInJSON BeerInputJSON
	err := json.Unmarshal(c.Body(), &beerInJSON)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	validate := validator.New()
	err = validate.Struct(beerInJSON)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return c.SendStatus(http.StatusBadRequest)
		}

		for _, err := range err.(validator.ValidationErrors) {
			msg := ErrorResponseJSON{
				Message: fmt.Sprintf("field %s invalid", err.Field()),
			}

			return c.Status(http.StatusBadRequest).JSON(msg)
		}
	}

	c.Locals(keyInput, beerInJSON)

	return c.Next()
}
