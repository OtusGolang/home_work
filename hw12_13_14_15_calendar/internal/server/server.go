package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server interface {
	Start() error
	Stop() error
}

type ServerInstance struct {
	instance *http.Server
}

type BasicHandler func(http.ResponseWriter, *http.Request)

func logMiddleware(h BasicHandler) BasicHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(t time.Time) {
			log.Println(r.RemoteAddr+" "+r.Method+" "+r.Host+" "+r.UserAgent(), " ", time.Since(t).Milliseconds(), "ms")
		}(time.Now())

		h(w, r)
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func (s *ServerInstance) Start() error {
	s.instance = &http.Server{Addr: ":8080"}
	http.HandleFunc("/hello", logMiddleware(helloHandler))
	err := s.instance.ListenAndServe()
	if err != nil {
		return err
	}
	fmt.Println("server started at port :8080")
	return nil
}

func (s *ServerInstance) Stop() error {
	return s.instance.Shutdown(context.Background())
}
