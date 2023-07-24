package application

import (
	"encoding/json"
	"home-service/pkg/application/entities"
	"home-service/pkg/infra/home"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHome(r *mux.Router) {
	h, err := home.InitialiseHomeProvider()
	if err != nil {
		log.Printf("error starting home, err %v", err)
	}
	HomeProvider = h

	r.HandleFunc("/test/{text}", testHandler).Methods(http.MethodGet)
	r.HandleFunc("/endpoint/{name}", addEndpointHandler).Methods(http.MethodPost)
	r.HandleFunc("/condition/{name}", addConditionHandler).Methods(http.MethodPost)
	r.HandleFunc("/restart", restartHandler).Methods(http.MethodPost)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	n := params[name]

	body := entities.NewResponse(n)

	writeReponse(w, r, body)
}

func addEndpointHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	n := params[name]

	var inBody entities.ReqIn
	err := json.NewDecoder(r.Body).Decode(&inBody)
	if err != nil {
		log.Printf("error reading body, err %v", err)
	}

	err = HomeProvider.AddRoute(n, inBody.Route, inBody.Method)
	if err != nil {
		log.Printf("error adding route, err %v", err)
	}

	body := entities.NewResponse("endpoint added", err)

	writeReponse(w, r, body)
}

func addConditionHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	n := params[name]

	var inBody entities.ReqIn
	err := json.NewDecoder(r.Body).Decode(&inBody)
	if err != nil {
		log.Printf("error reading body, err %v", err)
	}

	err = HomeProvider.AddCondition(n, inBody.Condition)
	if err != nil {
		log.Printf("error adding condition, err %v", err)
	}

	body := entities.NewResponse("condition added", err)

	writeReponse(w, r, body)
}

func restartHandler(w http.ResponseWriter, r *http.Request) {

	errStop := HomeProvider.StopHome()
	if errStop != nil {
		log.Printf("error stopping home, err %v", errStop)
	}

	errStart := HomeProvider.StartHome()
	if errStart != nil {
		log.Printf("error starting home, err %v", errStart)
	}

	body := entities.NewResponse("service restarted", errStop, errStart)

	writeReponse(w, r, body)
}
