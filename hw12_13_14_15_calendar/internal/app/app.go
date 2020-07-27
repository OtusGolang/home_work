package app

import "calendar/internal/repository"

type Repository interface {
	GetEvents() error
}

type App struct {
	r repository.BaseRepo
}

func New(r Repository) (*App, error) {
	return &App{r: r}, nil
}

func (a *App) Run() error {
	return nil
}