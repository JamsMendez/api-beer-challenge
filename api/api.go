package api

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"api-beer-challenge/internal/service"
)

const port = ":3000"

type API struct {
	service service.Service
	app     *fiber.App
}

func (a *API) Start() {
	rHandler := newRouterHandler(a.service)
	setUpRouters(a.app, rHandler)

	err := a.app.Listen(port)
	log.Fatal(err)
}

func New(s service.Service) *API {
	return &API{
		s,
		fiber.New(),
	}
}

func setUpRouters(app *fiber.App, rHandler *routerHandler) {
	groupRouter := app.Group(prefixAPI)
	groupRouter.Get("/beers", rHandler.getBeers)
	groupRouter.Get("/beers/:id", rHandler.getBeer)
	groupRouter.Get("/beers/:beerID/boxprice", rHandler.getBeerBoxPrice)
	groupRouter.Post("/beers", rHandler.addBeer)
}
