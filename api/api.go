package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"api-beer-challenge/internal/service"
)

const (
	prefixAPI = "/api"
)

type API struct {
	service service.Service
	app     *fiber.App
}

func (a *API) Start(port int) {
	rHandler := NewRouterHandler(a.service)
	SetUpRouters(a.app, rHandler)

	nPort := fmt.Sprintf(":%d", port)
	err := a.app.Listen(nPort)
	log.Fatal(err)
}

func New(s service.Service) *API {
	return &API{
		s,
		fiber.New(),
	}
}

func SetUpRouters(app *fiber.App, rHandler *routerHandler) {
	groupRouter := app.Group(prefixAPI)
	groupRouter.Get("/beers", rHandler.getBeers)
	groupRouter.Get("/beers/:id", validateReqGetBeer, rHandler.getBeer)
	groupRouter.Get("/beers/:beerID/boxprice", validateReqGetBeerBoxPrice, rHandler.getBeerBoxPrice)
	groupRouter.Post("/beers", validateReqAddBeer, rHandler.addBeer)
}
