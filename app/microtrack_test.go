package microtrack

import (
	"context"
	"log"
	"testing"

	"github.com/tidwall/gjson"
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

func Test_GetAPIToken(t *testing.T) {
	s := GetApiToken(
		"gengelmann",
		"RcTybM3hcV",
		"de7c21f6-fb85-49cb-b6a4-f80925d7de74",
	)
	log.Printf("%v", s)
	value := gjson.Get(s, "access_token")
	log.Printf("%v", value.String())

}

func Test_processResumenMensual(t *testing.T) {
	IDPlanilla := "e6da90a7853f4aafa1ae8308c5ebaa7f"
	raw := getResumenMensual(202001, IDPlanilla, 1)

	log.Printf("%v", processResumenMensual(raw, 202001))

}
