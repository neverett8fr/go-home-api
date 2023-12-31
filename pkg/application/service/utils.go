package application

import (
	"encoding/json"
	"home-service/pkg/infra/home"
	"log"
	"net/http"
)

const (
	name = "name"
)

var (
	HomeProvider home.HomeProvider
)

func writeReponse(w http.ResponseWriter, r *http.Request, body interface{}) {

	reponseBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("error converting reponse to bytes, err %v", err)
	}
	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(reponseBody)
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}
