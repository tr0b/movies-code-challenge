package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/tr0b/movies-code-challenge/backend/api"
	db "github.com/tr0b/movies-code-challenge/backend/db/sqlc"
	"github.com/tr0b/movies-code-challenge/backend/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	log.Println("Loaded config successfully")

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	log.Println("Connected successfully to db")

	store := db.NewStore(conn)

	log.Println("Created db store successfully")

	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	log.Println("created server successfully")
	log.Println("Server: ", server)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
