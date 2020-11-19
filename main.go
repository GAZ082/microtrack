package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	microtrack "github.com/gaz082/microtrack/app"
)

func main() {

	host := flag.String("host", "localhost", "host")
	port := flag.Int("port", 5432, "port")
	database := flag.String("database", "", "database")
	user := flag.String("user", "", "user")
	pass := flag.String("pass", "", "password")
	schema := flag.String("schema", "", "table schema")
	whatdata := flag.String("whatdata", "", "what data to retrieve [monthBrief (requires from, to, be in the same month, 1 to end of month, ie: 2020-01-01 2020-01-31)]")
	from := flag.String("from", "", "date from")
	to := flag.String("to", "", "date to")
	flag.Parse()
	if *database == "" || *user == "" || *pass == "" || *schema == "" || *whatdata == "" || *from == "" || *to == "" {
		log.Printf(
			`MICROTRACK DATA LOAD (c) 2020 Gabriel A. Zorrilla - @GabrielZorrilla%v
			Build: %v%v
			Missing arguments%v
			Example: microtrack -database=database -user=user -pass="password" -whatdata=monthBrief -schema=public%v -from=2020-03-01 -to=2020-03-21
			Use microtrack -h to check default values. Log in /tmp/microtrack.log
		`, "\n", time.Now().Format(time.RFC3339), "\n", "\n", "\n")
		os.Exit(1)
	}

	microtrack.Run([]string{
		*host,
		strconv.Itoa(*port),
		*database,
		*user,
		*pass, //4
		*whatdata,
		*schema,
		*from,
		*to})
}
