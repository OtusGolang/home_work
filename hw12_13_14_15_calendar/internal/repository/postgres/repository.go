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
	r.db, err = sqlx.Connect("pgx", "host=localhost port=5432 user=yanis password=yanis dbname=events sslmode=disable")
	return
}

func (r *Repo) Close() error {
	return r.db.Close()
}

func (r *Repo) AddEvent(event repository.Event) (err error) {
	var events []repository.Event

	nstmt, err := r.db.PrepareNamed(
		"INSERT INTO events (title, start_at, end_at, description, user_id, notify_at) VALUES (:title, :start_at, :end_at, :description, :user_id, :notify_at)")

	if err != nil {
		return
	}

	err = nstmt.Select(&events, event)

	return err
}

func (r *Repo) UpdateEvent(event repository.Event) (err error) {
	var events []repository.Event

	nstmt, err := r.db.PrepareNamed(
		"UPDATE events SET title=:title, start_at=:start_at, end_at = :end_at, description = :description, notify_at=:notify_at WHERE  user_id = :user_id and id=:id")

	if err != nil {
		return
	}

	err = nstmt.Select(&events, event)

	return
}

func (r *Repo) DeleteEvent(userId repository.Id, eventId repository.Id) (err error) {
	var events []repository.Event
	option := make(map[string]interface{})
	option["event_id"] = eventId
	option["user_id"] = userId

	nstmt, err := r.db.PrepareNamed("DELETE FROM events WHERE  user_id = :user_id and id=:event_id")

	if err != nil {
		return
	}

	err = nstmt.Select(&events, option)

	return
}

func (r *Repo) getEvents(userId repository.Id, from time.Time, to time.Time) ([]repository.Event, error) {
	var events []repository.Event
	option := make(map[string]interface{})
	option["start"] = from
	option["end"] = to
	option["user_id"] = userId

	nstmt, err := r.db.PrepareNamed("SELECT * FROM events WHERE  user_id = :user_id and start_at>=:start and start_at<:end")

	if err != nil {
		return nil, err
	}

	err = nstmt.Select(&events, option)

	return events, err
}

func (r *Repo) GetEventsDay(userId repository.Id, from time.Time) ([]repository.Event, error) {
	return r.getEvents(userId, from, from.Add(time.Hour*24))
}

func (r *Repo) GetEventsWeek(userId repository.Id, from time.Time) ([]repository.Event, error) {
	return r.getEvents(userId, from, from.AddDate(0, 0, 7))
}

func (r *Repo) GetEventsMonth(userId repository.Id, from time.Time) ([]repository.Event, error) {
	return r.getEvents(userId, from, from.AddDate(0, 1, 0))
}
