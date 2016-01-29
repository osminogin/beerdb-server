package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

const dbType = "mysql"

func init() {
	dsn := strings.TrimPrefix(os.Getenv("CLEARDB_DATABASE_URL"), dbType+"://")
	db, _ = sql.Open(dbType, dsn)
	if err := db.Ping(); err != nil {
		log.Fatal("Can't connect to database")
		os.Exit(1)
	}
}

func main() {
	port := os.Getenv("PORT")
	api := mux.NewRouter().StrictSlash(true)

	api.HandleFunc("beer", GetBeerListHandler(db))
	api.HandleFunc("beer/{key}", GetBeerByKeyHandler(db))
	api.HandleFunc("brand", GetBrandListHandler(db))
	api.HandleFunc("brand/{key}", GetBrandByKeyHandler(db))
	api.HandleFunc("brewery", GetBreweryListHandler(db))
	api.HandleFunc("brewery/{key}", GetBreweryKeyHandler(db))

	http.Handle("/api", api)
	http.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
