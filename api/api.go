package api

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"api-beer-challenge/internal/service"
)

const (
	port      = ":3000"
	prefixAPI = "/api"
)

type API struct {
	service service.Service
	app     *fiber.App
}

func (a *API) Start() {
	rHandler := NewRouterHandler(a.service)
	SetUpRouters(a.app, rHandler)

	err := a.app.Listen(port)
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
