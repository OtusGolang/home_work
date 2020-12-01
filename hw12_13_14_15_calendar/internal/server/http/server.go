package internalhttp

import (
	"context"
)

type Server struct {
	// TODO
}

type Application interface {
	// TODO
}

func NewServer(app Application) *Server {
	return &Server{}
}

func (s *Server) Start(ctx context.Context) error {
	// TODO
	select {
	case <-ctx.Done():
		return nil
	}
}

func (s *Server) Stop(ctx context.Context) error {
	// TODO
	return nil
}

// TODO
