package handler

import (
	"fmt"
	"github.com/Madq92/etcd-http-client/etcd"
	"github.com/Madq92/etcd-http-client/log"
	"github.com/gorilla/mux"
	"net/http"
)

func GetKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	key := vars["key"]
	value, err := etcd.GetKey(key)
	if err != nil {
		log.LOGGER.Error("error")
	}

	fmt.Fprintf(w, "%s = %$", key, value)
}

func PutKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	params := r.URL.Query()

	key := params.Get("key")
	value := params.Get("value")
	version, err := etcd.PutKey(key, value)
	if err != nil {
		log.LOGGER.Error("error")
	}

	fmt.Fprintf(w, "version = %n", version)
}
