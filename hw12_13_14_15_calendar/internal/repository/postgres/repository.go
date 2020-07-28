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
	return []repository.Event{}, nil
}