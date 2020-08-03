package main

import (
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
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
	//args := getArgs()
	//
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	////c, _ := config.Read("/home/yanis/work/home_work/hw12_13_14_15_calendar/configs/local.toml")
	//c, _ := config.Read(args.configPath)
	//fmt.Println("%+v", c)
	//
	//r := new(memory.MemoryDb)
	//s := new(server.ServerInstance)
	//l := new(logger.LoggerInstance)
	////if err := r.Connect(ctx, c.PSQL.DSN); err != nil {
	////	log.Fatal(err)
	////}
	////defer r.Close()
	//
	//a, err := app.New(r, s, l)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//cancel()
	//if err := a.Run(ctx); err != nil {
	//	log.Fatal(err)
	//}

	db, err := sqlx.Connect("pgx", "user=yanis password=yanis dbname=events sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv

	// Query the database, storing results in a []Person (wrapped in []interface{})
	events := []Event{}
	db.Select(&events, "SELECT * FROM events")
	//jason, john := events[0], events[1]

	for _, event := range events {
		fmt.Println("%v", event)
	}
}
