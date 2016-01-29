package main

import (
	"database/sql"
	_ "github.com/osminogin/beerdb-server/Godeps/_workspace/src/github.com/go-sql-driver/mysql"
	"github.com/osminogin/beerdb-server/Godeps/_workspace/src/github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
	"regexp"
)

var db *sql.DB

const dbType = "mysql"

func init() {
	re, _ := regexp.Compile(`@(.*)/`)
	log.Printf("AAA: %v", re.FindStringSubmatch(os.Getenv("CLEARDB_DATABASE_URL")))
	dsn := strings.TrimPrefix(os.Getenv("CLEARDB_DATABASE_URL"), dbType+"://")
	db, _ = sql.Open(dbType, dsn)
	if err := db.Ping(); err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	api := mux.NewRouter().StrictSlash(true)

	api.Handle("beer", GetBeerList(db))
	api.Handle("beer/{key}", GetBeerByKey(db))
	api.Handle("brand", GetBrandList(db))
	api.Handle("brand/{key}", GetBrandByKey(db))
	api.Handle("brewery", GetBreweryList(db))
	api.Handle("brewery/{key}", GetBreweryKey(db))

	http.Handle("/api", api)
	http.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
