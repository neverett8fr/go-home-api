package application

import (
	"home-service/pkg/application/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHome(r *mux.Router) {
	r.HandleFunc("/test/{text}", testHandler).Methods("GET")
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := entities.NewResponse(text)

	writeReponse(w, r, body)
}
