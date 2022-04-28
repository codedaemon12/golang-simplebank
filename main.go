package main

import (
	"database/sql"
	"gopractice/simplebank/api"
	db "gopractice/simplebank/db/sqlc"
	"log"
)

const dbDriver = "postgres"
const dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
const serverAddress = "0.0.0.0:8000"

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connnec to db : ", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
