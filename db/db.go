package db

import (
	"context"
	"fmt"
	mt "gomicrotrack/microtrack"
	"gomicrotrack/tool"
	"log"

	"github.com/jackc/pgx/v4"
	database "github.com/jackc/pgx/v4"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	//host     = "192.168.0.224"
	port     = 5432
	user     = "leer_insertar_user"
	password = "taza blanca cocinero"
	dbname   = "microtrack"
)

// db is the global database connection.
var db, _ = connect()

// Connect DB connection.
func connect() (*database.Conn, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password='%s' dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := database.Connect(context.Background(), connString)
	return db, err
}

func doQuery(records []*mt.DatosFranja) (recordsQuantity int) {
	rows := [][]interface{}{}
	for _, row := range records {
		rows = append(rows, []interface{}{row.Fecha.Time, row.DuracionSegundos, row.Latitud, row.Longitud, row.Metros, row.Operador, row.OperadorID,
			row.Sitio, row.TipoEvento, row.Vehiculo, row.Zona})
	}
	_, err := db.CopyFrom(context.Background(),
		pgx.Identifier{"data"}, []string{"timestamp", "duration", "lat", "lon", "meters", "operator_description", "operator_id", "site", "event_type", "vehicle_id", "zone"}, pgx.CopyFromRows(rows))
	tool.CheckError(err, "info")
	return len(records)
}

func InsertRecords(records []*mt.DatosFranja) {
	txn, err := db.Begin(context.Background()) // db is never closed
	tool.CheckError(err, "i")
	q := doQuery(records)
	err = txn.Commit(context.Background())
	if err != nil {
		log.Fatal("There was a problem. Nothing inserted.")
	} else {
		log.Printf("Inserted %v records.", q)
	}

}
