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

// InitDB initialize and check connection to database.
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

	db, _ = gorm.Open(dbDriver, dbSource)

	// Check database connection
	if err := db.DB().Ping(); err != nil {
		return err
	}
	return nil
}
