package main

import (
	"calendar/internal/app"
	"calendar/internal/config"
	"fmt"
	"log"
)

func main() {
	c, _ := config.Read("/home/yanis/work/home_work/hw12_13_14_15_calendar/configs/local.toml")
	fmt.Println("%+v", c)

	a, err := app.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}

