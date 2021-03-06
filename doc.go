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
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/osminogin/beerdb-server/Godeps/_workspace/src/github.com/russross/blackfriday"
)

// ReadmeHandler read README.md file from parent directory and include
// rendered html to base template.
func ReadmeHandler(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}
	output := blackfriday.MarkdownCommon(input)
	t, _ := template.ParseFiles("templates/base.html")
	t.Execute(w, output)
}
