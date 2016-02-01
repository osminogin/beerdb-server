package main

import (
	"os"
	"fmt"
	"regexp"
	"github.com/osminogin/beerdb-server/Godeps/_workspace/src/github.com/jinzhu/gorm"
	_ "github.com/osminogin/beerdb-server/Godeps/_workspace/src/github.com/go-sql-driver/mysql"
)

var (
	db gorm.DB
)

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