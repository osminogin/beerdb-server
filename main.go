package main

import (
	"log"
	"net/http"
	"os"

	"github.com/osminogin/beerdb-server/Godeps/_workspace/src/github.com/gorilla/mux"
)

func main() {
	// Database initialization
	err := InitDB()
	if err != nil {
		log.Fatalf("DB error: %v", err)
		os.Exit(1)
	}
	defer db.DB().Close()

	// HTTP routes
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

	// Start server
	port := os.Getenv("PORT")
	log.Printf("Starting beerdb-server on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
