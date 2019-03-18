package handler

import (
	"encoding/json"
	"github.com/Madq92/etcd-http-client/config"
	"net/http"
)

func GetConf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(config.Conf)
}
