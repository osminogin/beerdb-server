package main

import (
	"os"
	"log"
	"net/http"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
