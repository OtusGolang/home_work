package postgres

import (
	"calendar/internal/repository"
	"context"
)

var _ repository.BaseRepo = (*Repo)(nil)

func (r Repo) Connect(ctx context.Context) error {
	panic("implement me")
}

func (r Repo) Close() error {
	panic("implement me")
}

func (r Repo) GetBooks(ctx context.Context) ([]repository.Event, error) {
	panic("implement me")
}

type Repo struct {

}