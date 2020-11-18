package main

import (
	"gomicrotrack/db"
	mt "gomicrotrack/microtrack"

	"log"
	"os"
)

func start(args []string) {

	log.Print(`
	=== MICROTRACK DATA LOAD ===
	V0.0 2020-03-20
	gabriel.zorrilla@galileoar.com
	Usage: microtrackload from to
	from, to:  dates with format YYYY-MM-DD
`)

	//the first argument is the command
	switch len(args) - 1 {
	case 1:
		log.Printf("Inserting records to database of %v", args[1])
		db.InsertRecords(mt.DayDetail(args[1], args[1]))
	case 2:
		log.Printf("Inserting records to database from %v to %v", args[1], args[2])
		db.InsertRecords(mt.DayDetail(args[1], args[2]))
	default:
		log.Print("Missing parameters.")
	}
}

func main() {
	start(os.Args)
}
