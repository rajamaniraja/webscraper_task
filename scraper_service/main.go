package main

import (
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
	handlers "github.com/rajamaniraja/webscraper_task/scraper_service/handlers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/healthz", handlers.HealthHandler)

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	subrouter.HandleFunc("/scrap", handlers.ScrapUrl).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
