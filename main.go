package main

import (
	r "./router"
	"github.com/Madq92/etcd-http-client/log"
	"net/http"
)

func main() {
	router := r.NewRouter()
	log.LOGGER.Info("Server started")
	log.LOGGER.Fatal(http.ListenAndServe(":8080", router))
}
