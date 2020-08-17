package internalhttp

import "context"

type Server struct {
	// TODO
}

func NewServer(a AppI) *Server {
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
