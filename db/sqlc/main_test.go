package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecretpassword@localhost:5432/bank-app?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("cannot connect to DB: %s", err.Error())
		return
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
