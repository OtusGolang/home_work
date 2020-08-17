package app

import (
	"context"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	// TODO
}

func New(logger LoggerI, storage StorageI) *App {
	return &App{}
}

func (a *App) CreateEvent(ctx context.Context, id string, title string) error {
	return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
