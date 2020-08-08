package main

import (
	"calendar/internal/app"
	"calendar/internal/config"
	"calendar/internal/logger"
	"calendar/internal/repository/postgres"
	"calendar/internal/server"
	"context"
	"flag"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

type Args struct {
	configPath string
}

func getArgs() *Args {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	args := Args{
		configPath: *configPath,
	}

	return &args
}

func main() {
	args := getArgs()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// ./configs/local.toml
	c, _ := config.Read(args.configPath)

	r := new(postgres.Repo)
	s := new(server.ServerInstance)
	l := new(logger.LoggerInstance)

	a, err := app.New(r, s, l)
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Run(ctx, c.Logger.Path, c.PSQL.DSN); err != nil {
		log.Fatal(err)
	}
}
