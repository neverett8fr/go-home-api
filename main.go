package main

import (
	"fmt"
	cmd "home-service/cmd"
	application "home-service/pkg/application/service"
	"home-service/pkg/config"
	"log"

	"github.com/gorilla/mux"
)

// Route declaration
func getRoutes() *mux.Router {
	r := mux.NewRouter()
	application.NewHome(r)

	return r
}

// Initiate web server
func main() {

	conf := config.Initialise()

	fmt.Println("service started")
	router := getRoutes()

	err := cmd.StartServer(*conf, router)
	if err != nil {
		log.Fatalf("error starting server, %v", err)
	}

}
