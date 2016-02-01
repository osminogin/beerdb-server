package main

import (
	"os"
	"log"
	"net/http"

	"github.com/osminogin/beerdb-server/Godeps/_workspace/src/github.com/gorilla/mux"
)

func init() {
	// Database initialization
	err := InitDB()
	if err != nil {
		log.Fatalf("DB error: %v", err)
		os.Exit(1)
	}
	// TODO: Defer database connection close
}

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/", ReadmeHandler).Methods("GET")

	api := r.PathPrefix("/api/v1").Subrouter()
	api.Handle("/beer", GetBeerList(&db))
	api.Handle("/beer/{key}", GetBeerByKey(&db))
	api.Handle("/brand", GetBrandList(&db))
	api.Handle("/brand/{key}", GetBrandByKey(&db))
	api.Handle("/brewery", GetBreweryList(&db))
	api.Handle("/brewery/{key}", GetBreweryKey(&db))

	http.Handle("/", r)
	log.Printf("Starting beerdb-server on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
