package main

import (
	"database/sql"
	"log"
	"simple-bank/api"
	db "simple-bank/db/sqlc"
	"simple-bank/util"

	_ "github.com/lib/pq"
)


func main(){
	//loading the configuration files
	config, configError := util.LoadConfig(".");
	if configError != nil {
		log.Fatal("Error in configuration Loading", configError)
	}
	
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)

	//starting the server

	server := api.NewServer(store)

	server.Start(config.ServerAddress)
}