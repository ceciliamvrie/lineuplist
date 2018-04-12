package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/techmexdev/lineuplist/pkg/handler"
	"github.com/techmexdev/lineuplist/pkg/storage/postgres"
)

func main() {
	goEnv := os.Getenv("GO_ENV")
	dsn := os.Getenv("PG_DSN")

	var options handler.Options
	if goEnv == "PROD" {
		options = handler.Options{Log: false}
	} else {
		options = handler.Options{Log: true}
	}

	router := handler.New(postgres.New(dsn), options)

	if goEnv == "PROD" {
		log.Println("Starting server at port 80...")
		http.ListenAndServeTLS(":80", "server.crt", "server.key", router)
	} else {
		log.Println("Starting server at localhost:3000...")
		log.Fatal(http.ListenAndServe(":3000", router))
	}
}