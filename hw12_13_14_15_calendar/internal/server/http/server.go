package internalhttp

import "context"

type Server struct {
	// TODO
}

type Application interface {
	// TODO
}

func NewServer(app Application) *Server {
	return &Server{}
}

func (s *Server) Start() error {
	// TODO
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	// TODO
}

// TODO
