// Copyright 2016 Vladimir Osintsev <osintsev@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
// See the LICENSE file in the main directory for details.

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/osminogin/beerdb-server/Godeps/_workspace/src/github.com/gorilla/mux"
)

func main() {
	// Database initialization
	if err := InitDB(); err != nil {
		log.Fatalf("DB error: %v", err)
	}
	defer db.DB().Close()

	// HTTP routes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ReadmeHandler).Methods("GET")

	api := router.PathPrefix("/api/v1").Subrouter()
	api.Handle("/beer/", GetBeerList(&db))
	api.Handle("/beer/{key}", GetBeerByKey(&db))
	api.Handle("/brand/", GetBrandList(&db))
	api.Handle("/brand/{key}", GetBrandByKey(&db))
	api.Handle("/brewery/", GetBreweryList(&db))
	api.Handle("/brewery/{key}", GetBreweryKey(&db))

	// Start server
	port := os.Getenv("PORT")
	log.Printf("Starting beerdb-server on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
