package main

import (
	"calendar/internal/app"
	"calendar/internal/config"
	"calendar/internal/repository/postgres"
	"context"
	"fmt"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, _ := config.Read("/home/yanis/work/home_work/hw12_13_14_15_calendar/configs/local.toml")
	fmt.Println("%+v", c)

	r := new(postgres.Repo)
	if err := r.Connect(ctx, c.PSQL.DSN); err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	a, err := app.New(r)
	if err != nil {
		log.Fatal(err)
	}

	cancel()
	if err := a.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

