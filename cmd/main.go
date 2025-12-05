package main

import (
	"context"
	"log"

	db "github.com/amrrdev/refx/db"
	"github.com/amrrdev/refx/internal/app"
	"github.com/amrrdev/refx/internal/config"
	"github.com/amrrdev/refx/internal/database"
	"github.com/amrrdev/refx/internal/redis"
	"github.com/amrrdev/refx/internal/url"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalln("could not load env:", err)
	}

	err = db.RunMigrations(config.DatabaseUrl)
	if err != nil {
		log.Fatalln(err)
	}

	context := context.Background()
	database, err := database.NewDatabase(context, config.DatabaseUrl)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := redis.NewClient(config.RedisConnection)
	if err != nil {
		log.Fatalln(err)
	}

	repo := url.NewRepository(database)
	service := url.NewService(repo, client)
	handler := url.NewHandler(service)
	r := app.NewServer(handler)

	r.Run()
	defer database.Pool.Close()
}
