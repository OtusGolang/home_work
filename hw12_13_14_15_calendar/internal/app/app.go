package app

import (
	"calendar/internal/logger"
	"calendar/internal/repository"
	"calendar/internal/server"
	"context"
	"log"
)

type App struct {
	repo repository.BaseRepo
	server server.Server
	logger logger.Logger
}

func New(r repository.BaseRepo, s server.Server, l logger.Logger) (*App, error) {
	return &App{repo: r, server: s, logger: l}, nil
}

func (a *App) Run(ctx context.Context, logPath string, dsn string) error {
	// logger
	err := a.logger.Init(logPath)
	if err != nil {
		log.Fatal(err)
	}

	// server
	err = a.server.Start()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: don't forget to close connection
	// storage
	err = a.repo.Connect(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
