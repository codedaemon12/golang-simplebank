package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const dbDriver = "postgres"
const dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connnec to db : ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
