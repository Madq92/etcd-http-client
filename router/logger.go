package router

import (
	"../log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.LOGGER.Printf("%s %s %s %s", r.Method, r.RequestURI, name, time.Since(start))
	})
}
