package app

import (
	"calendar/internal/repository"
	"context"
	"fmt"
)

type App struct {
	r repository.BaseRepo
}

func New(r repository.BaseRepo) (*App, error) {
	return &App{r: r}, nil
}

func (a *App) Run(ctx context.Context) error {
	events, _ := a.r.GetEvents(ctx)

	for _, event := range events {
		fmt.Println("%v", event)
	}

	return nil
}