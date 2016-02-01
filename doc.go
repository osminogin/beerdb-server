package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"github.com/osminogin/beerdb-server/Godeps/_workspace/src/github.com/russross/blackfriday"
)

func ReadmeHandler(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}
	unsafe := blackfriday.MarkdownCommon(input)
	w.Write(unsafe)
}
