package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Server interface {
	Start() error
	Stop() error
}

type ServerInstance struct {
	instance *http.Server
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println(GetIP(req) + " " + req.Method + " " + req.Host + " " + req.UserAgent())
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
