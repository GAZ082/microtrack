package main

import (
	"flag"
	"log"
	"os"
	"time"

	microtrack "github.com/gaz082/microtrack/app"
)

func main() {
	whatdata := flag.String("whatdata", "", "data to retrieve. Options: [monthBrief (requires period)]")
	period := flag.String("period", "", "Period (usually month) of the request. Format YYYYMM, integer. Ie: 202012 (for december 2020).")
	templateID := flag.String("templateID", "", "Identifier of the MT template. User must have access to this template.")
	flag.Parse()
	if *whatdata == "" || *period == "" || *templateID == "" {
		log.Printf(
			`MICROTRACK DATA LOAD (c) 2020 Gabriel A. Zorrilla - @GabrielZorrilla%v
			Build: %v%v
			Missing arguments%v
			Example: microtrack -whatdata=monthBrief -period=202012 -templateID=XXXX
			Use microtrack -h to check default values. Log in /tmp/microtrack.log
		`, "\n", time.Now().Format(time.RFC3339), "\n", "\n")
		os.Exit(1)
	}

	microtrack.Run([]string{
		*whatdata,
		*period,
		*templateID,
	})
}
