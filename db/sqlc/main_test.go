package db

import (
	"database/sql"
	"log"
	"os"
	"simple-bank/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, configError := util.LoadConfig("../../");
	if configError != nil {
		log.Fatal("Error in configuration Loading", configError)
	}
	var err error
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
