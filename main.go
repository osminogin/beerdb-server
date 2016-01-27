package main

import (
	"os"
	"log"
	"fmt"
	"net/http"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello, heroku")
}