package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQuery *Queries
var testDB *sql.DB

const (
	driverName     = "postgres"
	dataSourceName = "postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}

	testQuery = New(testDB)

	os.Exit(m.Run())
}
