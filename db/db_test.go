package db

import (
	mt "gomicrotrack/microtrack"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func Test_connect(t *testing.T) {
	_, err := connect()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("All good.")
	}
}

func Test_InsertRecords(t *testing.T) {
	InsertRecords(mt.DayDetail("2020-01-02", "2020-01-03"))
}
