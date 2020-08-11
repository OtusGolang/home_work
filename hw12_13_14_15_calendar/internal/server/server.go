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

type Instance struct {
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

func (s *Instance) Start() error {
	s.instance = &http.Server{Addr: ":8080"}
	http.HandleFunc("/hello", logMiddleware(helloHandler))
	fmt.Println("server starting at port :8080")
	return s.instance.ListenAndServe()
}

func (s *Instance) Stop() error {
	return s.instance.Shutdown(context.Background())
}
