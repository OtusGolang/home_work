package postgres

import (
	"calendar/internal/repository"
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

type Repo struct {
	db *sqlx.DB
}

func (r *Repo) Connect(ctx context.Context, dsn string) (err error) {
	//r.db, err = sql.Open("pgx", dsn)
	//if err != nil {
	//	return
	//}
	//// TODO: WTF?
	//return r.db.PingContext(ctx)
	fmt.Println("inside connect")
	r.db, err = sqlx.Connect("pgx", "user=yanis password=yanis dbname=events sslmode=disable")
	return
}

func (r *Repo) Close() error {
	return r.db.Close()
}

func (m *Repo) AddEvent(event repository.Event) error {
	panic("123")
}

func (m *Repo) UpdateEvent(event repository.Event) error {
	panic("123")
}

func (m *Repo) DeleteEvent(userId repository.Id, eventId repository.Id) error {
	panic("123")
}


func (r *Repo) GetEventsDay(userId repository.Id, from time.Time) ([]repository.Event, error) {
	var events []repository.Event
	err := r.db.Select(&events, "SELECT * FROM events")

	return events, err
}

func (m *Repo) GetEventsWeek(userId repository.Id, from time.Time) ([]repository.Event, error) {
	panic("123")
}

func (m *Repo) GetEventsMonth(userId repository.Id, from time.Time) ([]repository.Event, error) {
	panic("123")
}
