package server

import (
	"context"
	"fmt"
	"net/http"
)

type Server interface {
	Start() error
	Stop() error
}

type ServerInstance struct {
	instance *http.Server
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func (s *ServerInstance) Start() error {
	s.instance = &http.Server{Addr: ":8080"}
	http.HandleFunc("/hello", hello)

	return s.instance.ListenAndServe()
}

func (s *ServerInstance) Stop() error {
	return s.instance.Shutdown(context.Background())
}
