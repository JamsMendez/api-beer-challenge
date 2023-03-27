package main

import (
	"context"
	"log"

	"api-beer-challenge/api"
	"api-beer-challenge/database"
	"api-beer-challenge/internal/repository"
	"api-beer-challenge/internal/service"
	"api-beer-challenge/settings"
)

func main() {
	s, err := settings.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	db, err := database.GetConnection(ctx, s)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db)
	serv := service.New(repo)

	api.New(serv).Start()
}
