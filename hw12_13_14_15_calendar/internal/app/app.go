package app

import (
	"calendar/internal/logger"
	"calendar/internal/repository"
	"calendar/internal/server"
	"context"
	"log"
)

type App struct {
	r repository.BaseRepo
	s server.Server
	l logger.Logger
}

func New(r repository.BaseRepo, s server.Server, l logger.Logger) (*App, error) {
	return &App{r: r, s: s, l: l}, nil
}

func (a *App) Run(ctx context.Context) error {
	//events, err := a.r.GetEvents(ctx)
	path := "./logs/logs.txt"

	err := a.l.Init(path)
	if err != nil {
		log.Fatal(err)
	}

	err = a.s.Start()

	if err != nil {
		log.Fatal(err)
	}
	return nil
}
