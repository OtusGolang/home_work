package app

import (
	"context"
)

type App struct {
	// TODO
}

type Logger interface {
	// TODO
}

type Storage interface {
	// TODO
}

func New(logger Logger, storage Storage) *App {
	return &App{}
}

func (a *App) CreateEvent(ctx context.Context, id string, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
