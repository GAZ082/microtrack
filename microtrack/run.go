package microtrack

import (
	"context"
	"fmt"

	database "github.com/jackc/pgx/v4"
)

// Connect DB connection.
func connect(args []string) (*database.Conn, error) {
	//hostBD,puertoBD,nombreBD,usuarioBD,ppasswordBD,origen
	connString := fmt.Sprintf("host=%s port=%v dbname=%s user=%s password='%s'  sslmode=disable", args[0], args[1], args[2], args[3], args[4])
	db, err := database.Connect(context.Background(), connString)
	return db, err
}

func Run(args []string) {
	var err error
	DB, err = connect(args)
	checkError(err, "panic")
	if args[5] == "monthBrief" {
		InsertRecordsMontlhy(MonthlyBrief(args[7], args[8]))
	}
}