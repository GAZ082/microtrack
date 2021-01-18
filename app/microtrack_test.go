package microtrack

import (
	"context"
	"log"
	"testing"

	"github.com/tidwall/gjson"
)

func TestRun(t *testing.T) {
	args := []string{"monthBrief", "202004", "e6da90a7853f4aafa1ae8308c5ebaa7f"}
	Run(args)
}

func Test_connection(t *testing.T) {
	DB, err := connect()
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
		"xx",
		"xx",
		"xx",
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
