package main

import (
	"fmt"
	"net/http"
	"database/sql"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, heroku")
}

func GetBeerListHandler(db *sql.DB) http.Handler {
	return http.Handler(func(w http.ResponseWriter, r *http.Request) {

	})
}

func GetBeerByKeyHandler(db *sql.DB) http.Handler {
	return http.Handler(func(w http.ResponseWriter, r *http.Request) {

	})
}

func GetBrandListHandler(db *sql.DB) http.Handler {
	return http.Handler(func(w http.ResponseWriter, r *http.Request) {

	})
}

func GetBreweryKeyHandler(db *sql.DB) http.Handler {
	return http.Handler(func(w http.ResponseWriter, r *http.Request) {

	})
}

func GetBreweryListHandler(db *sql.DB) http.Handler {
	return http.Handler(func(w http.ResponseWriter, r *http.Request) {

	})
}

func GetBrandByKeyHandler(db *sql.DB) http.Handler {
	return http.Handler(func(w http.ResponseWriter, r *http.Request) {

	})
}
