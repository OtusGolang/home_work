package app

import (
	"context"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage"
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
	return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
