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
	"fmt"
	"os"
	"regexp"

	_ "github.com/osminogin/beerdb-server/Godeps/_workspace/src/github.com/go-sql-driver/mysql"
	"github.com/osminogin/beerdb-server/Godeps/_workspace/src/github.com/jinzhu/gorm"
)

var (
	db gorm.DB
)

// InitDB construct dbSource from environment and initialize
// connection to database with check.
func InitDB() error {
	databaseUrl := os.Getenv("CLEARDB_DATABASE_URL")
	re, _ := regexp.Compile("([^:]+)://([^:]+):([^@]+)@([^/]+)/([^?]+)")
	match := re.FindStringSubmatch(databaseUrl)

	dbDriver := match[1]
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?parseTime=true",
		match[2],
		match[3],
		match[4],
		match[5],
	)

	// Open connection to database
	db, _ = gorm.Open(dbDriver, dbSource)

	// Check database connection
	if err := db.DB().Ping(); err != nil {
		return err
	}
	return nil
}
