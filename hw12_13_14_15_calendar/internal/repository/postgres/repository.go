package postgres

import (
	"calendar/internal/repository"
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var _ repository.BaseRepo = (*Repo)(nil)

type Repo struct {
	db *sql.DB
}

func (r *Repo) Connect(ctx context.Context, dsn string) (err error) {
	r.db, err = sql.Open("pgx", dsn)
	if err != nil {
		return
	}
	return r.db.PingContext(ctx)
}

func (r *Repo) Close() error {
	return r.db.Close()
}

func (r *Repo) GetEvents(ctx context.Context) ([]repository.Event, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT title FROM post
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []repository.Event

	for rows.Next() {
		var e repository.Event
		if err := rows.Scan(&e.Title); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, rows.Err()
}