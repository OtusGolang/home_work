package main

import (
	"calendar/internal/config"
	"fmt"
)

func main() {
	c, _ := config.Read("/home/yanis/work/home_work/hw12_13_14_15_calendar/configs/local.toml")
	fmt.Println("%+v", c)
}

