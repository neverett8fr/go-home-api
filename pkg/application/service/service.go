package application

import (
	"home-service/pkg/application/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHome(r *mux.Router) {
	r.HandleFunc("/test/{text}", testHandler).Methods(http.MethodGet)
	r.HandleFunc("/endpoint/{name}", addEndpointHandler).Methods(http.MethodPost)
	r.HandleFunc("/condition/{name}", addConditionHandler).Methods(http.MethodPost)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := entities.NewResponse(text)

	writeReponse(w, r, body)
}

func addEndpointHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["name"]

	body := entities.NewResponse(text)

	writeReponse(w, r, body)
}

func addConditionHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["name"]

	body := entities.NewResponse(text)

	writeReponse(w, r, body)
}
