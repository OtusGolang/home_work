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

	//events, err := r.GetEventsDay(2, time.Now().Add(time.Hour * time.Duration(-5)))

	//err = r.DeleteEvent(1, 1)
	//err = r.AddEvent(repository.Event{
	//	Title:       "Event 6",
	//	StartAt:     time.Now(),
	//	EndAt:       time.Now(),
	//	Description: "Description 6",
	//	UserId:      1,
	//	NotifyAt:    time.Now(),
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//for _, event := range events {
	//	fmt.Println("%v", event)
	//}
}
