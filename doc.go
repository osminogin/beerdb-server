package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/osminogin/beerdb-server/Godeps/_workspace/src/github.com/russross/blackfriday"
)

// ReadmeHandler read README.md file from parent directory and
// include rendered html to template.
func ReadmeHandler(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}
	output := blackfriday.MarkdownCommon(input)
	t, _ := template.ParseFiles("templates/base.html")
	t.Execute(w, output)
}
