package main

import (
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
	handlers "github.com/rajamaniraja/webscraper_task/updater_service/handlers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/healthz", handlers.HealthHandler)

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	subrouter.HandleFunc("/store", handlers.StoreData).Methods("POST")

	log.Fatal(http.ListenAndServe(":80", router))
}
