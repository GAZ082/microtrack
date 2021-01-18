package microtrack

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	database "github.com/jackc/pgx/v4"
)

// db is the global database connection.
var DB *database.Conn

func doQueryMonthlyBrief(records *[]RecordResumenMensual) (recordsQuantity int) {
	rows := [][]interface{}{}
	for _, row := range *records {
		rows = append(rows, []interface{}{row.Period, row.AccCent, row.Accelerations, row.Braking, row.Identificador, row.InfHigh, row.InfLight, row.InfMedium, row.Kms, row.Score, row.ScoreAP, row.ScoreRR, row.ScoreZU})
	}
	_, err := DB.CopyFrom(context.Background(),
		pgx.Identifier{"resumen_mensual"}, []string{"period", "acccent", "accelerations", "braking", "identificador", "infhigh", "inflight", "infmedium", "km", "score", "scoreap", "scorerr", "scorezu"}, pgx.CopyFromRows(rows))
	checkError(err, "info")
	return len(*records)
}

func InsertRecordsMontlhy(records *[]RecordResumenMensual) {
	txn, err := DB.Begin(context.Background()) // db is never closed
	checkError(err, "i")
	q := doQueryMonthlyBrief(records)
	err = txn.Commit(context.Background())
	if err != nil {
		log.Fatal("There was a problem. Nothing inserted.")
	} else {
		log.Printf("Inserted %v records.", q)
	}
}
