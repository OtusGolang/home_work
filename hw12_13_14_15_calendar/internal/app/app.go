package app

import "calendar/internal/repository"

type App struct {
	r repository.BaseRepo
}

func New(r repository.BaseRepo) (*App, error) {
	return &App{r: r}, nil
}

func (a *App) Run() error {
	return nil
}