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
	instance *http.ServeMux
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

type BasicHandler func(http.ResponseWriter, *http.Request)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(t time.Time) {
			log.Println(GetIP(r)+" "+r.Method+" "+r.Host+" "+r.UserAgent(), " ", time.Since(t).Milliseconds(), "ms")
		}(time.Now())

		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func (s *ServerInstance) Start() error {
	siteMux := http.NewServeMux()
	siteMux.HandleFunc("/hello", helloHandler)

	siteHandler := logMiddleware(siteMux)

	fmt.Println("starting server at :8080")

	return http.ListenAndServe(":8080", siteHandler)
}

func (s *ServerInstance) Stop() error {
	return s.instance.Shutdown(context.Background())
}
