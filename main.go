package main

import (
	r "./router"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	//confInit()
	//logInit()

	log.Printf("Server started")
	router := r.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

//func confInit() {
//	configor.Load(&config.Config, "config.yml")
//}
//
//func logInit() {
//	log.SetOutput(os.Stdout)
//	log.SetLevel(log.WarnLevel)
//}
