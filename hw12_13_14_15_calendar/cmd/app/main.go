package main

import (
	"calendar/internal/app"
	"calendar/internal/config"
	"calendar/internal/repository/memory"
	"calendar/internal/server"
	"context"
	"flag"
	"fmt"
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

	r := new(memory.MemoryDb)
	s := new(server.ServerInstance)
	//if err := r.Connect(ctx, c.PSQL.DSN); err != nil {
	//	log.Fatal(err)
	//}
	//defer r.Close()

	a, err := app.New(r, s)
	if err != nil {
		log.Fatal(err)
	}

	cancel()
	if err := a.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
