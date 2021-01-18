package microtrack

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	database "github.com/jackc/pgx/v4"
)

// Connect DB connection.
func connect() (*database.Conn, error) {
	//hostBD,puertoBD,nombreBD,usuarioBD,ppasswordBD,origen
	connString := fmt.Sprintf("host=%s port=%v dbname=%s user=%s password='%s'  sslmode=disable", config.GetString("db.host"), config.GetString("db.port"),
		config.GetString("db.name"), config.GetString("db.user"), config.GetString("db.password"))
	db, err := database.Connect(context.Background(), connString)
	return db, err
}

func Run(args []string) {
	var logFile, _ = os.OpenFile("/tmp/microtrack.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(logFile)
	log.Print("----------::STARTING::----------")
	var err error
	DB, err = connect()
	checkError(err, "panic")
	if args[0] == "monthBrief" {
		period, _ := strconv.Atoi(args[1])
		raw := getResumenMensual(period, args[2], 1) //Use version 1 of the API.
		processed := processResumenMensual(raw, int32(period))
		InsertRecordsMontlhy(&processed)
	}
}
