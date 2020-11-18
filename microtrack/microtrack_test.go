package microtrack

import (
	"log"
	"testing"
)

func TestDayDetail(t *testing.T) {
	data := DayDetail("2020-03-25", "2020-03-25")
	log.Println(data[0])
	log.Println(data[len(data)-1])
}

func TestMonthlyBrief(t *testing.T) {
	result := MonthlyBrief("2020-08-01", "2020-08-01")

	for _, v := range result.ResumenMensualResult.DatosIdentificadores {
		log.Printf("%v", v)
	}

}
