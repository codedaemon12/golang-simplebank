package main

import (
	"database/sql"
	"gopractice/simplebank/api"
	db "gopractice/simplebank/db/sqlc"
	"gopractice/simplebank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connnec to db : ", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
