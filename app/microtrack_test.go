package microtrack

import (
	"context"
	"log"
	"testing"
)

var (
	args = []string{
		"45.33.2.159",
		"5432",
		"microtrack",
		"leer_insertar_user",
		"CLAVE",
		"monthBrief",
		"public",
		"2020-10-01",
		"2020-10-31",
	}
)

func TestDayDetail(t *testing.T) {
	data := DayDetail("2020-03-25", "2020-03-25")
	log.Printf("%+v\n", data[0])
}

func TestMonthlyBrief(t *testing.T) {
	result := MonthlyBrief("2020-10-01", "2020-10-31")
	for _, v := range result {
		log.Printf("%+v\n", v)
	}
}

func TestRun(t *testing.T) {
	Run(args)
}

func Test_connection(t *testing.T) {
	DB, err := connect(args)
	if err != nil {
		log.Print(err)
	}
	r, _ := DB.Query(context.Background(), "SELECT current_setting('server_version_num');")
	for r.Next() {
		var n string
		err = r.Scan(&n)
		log.Println(n)
	}
}
