package app

import (
	"calendar/internal/repository"
	"calendar/internal/server"
	"context"
	"log"
)

type App struct {
	r repository.BaseRepo
	s server.Server
}

func New(r repository.BaseRepo, s server.Server) (*App, error) {
	return &App{r: r, s: s}, nil
}

func (a *App) Run(ctx context.Context) error {
	//events, err := a.r.GetEvents(ctx)
	err := a.s.Start()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
