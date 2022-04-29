package db

import (
	"database/sql"
	"gopractice/simplebank/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err1 := util.LoadConfig("../..")
	if err1 != nil {
		log.Fatal("cannot read config ", err1)
	}
	var err error
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connnec to db : ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
