package main

import (
	"calendar/internal/app"
	"calendar/internal/config"
	"calendar/internal/logger"
	"calendar/internal/repository/postgres"
	"calendar/internal/server"
	"context"
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"time"
)

type Args struct {
	configPath string
}

func getArgs() *Args {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()
	//otherArgs := flag.Args()

	args := Args{
		configPath: *configPath,
	}

	return &args
}

type Event struct {
	Id          int
	Title       string
	StartAt     time.Time `db:"start_at"`
	EndAt       time.Time `db:"end_at"`
	Description string
	UserId      int       `db:"user_id"`
	NotifyAt    time.Time `db:"notify_at"`
}

func main() {
	args := getArgs()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//c, _ := config.Read("/home/yanis/work/home_work/hw12_13_14_15_calendar/configs/local.toml")
	c, _ := config.Read(args.configPath)
	fmt.Println("%+v", c)

	r := new(postgres.Repo)
	s := new(server.ServerInstance)
	l := new(logger.LoggerInstance)
	//if err := r.Connect(ctx, c.PSQL.DSN); err != nil {
	//	log.Fatal(err)
	//}
	//defer r.Close()

	fmt.Println("Hello0")

	a, err := app.New(r, s, l)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello1")

	//cancel()
	if err := a.Run(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello2")

	events, err := r.GetEventsDay(1, time.Now().Add(time.Hour * time.Duration(-5)))

	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		fmt.Println("%v", event)
	}
}
