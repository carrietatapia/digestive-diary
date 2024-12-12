package main

import (
	"log"

	"github.com/carrietatapia/digestivediary/internal/api/router"
	"github.com/carrietatapia/digestivediary/internal/api/server"
	"github.com/carrietatapia/digestivediary/internal/db"
)

func main() {
	var err error
	db.DB, err = db.ConnectDB("postgres://postgres:admin@localhost:5432/test")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.DB.Close()

	srv := server.NewServer(db.DB)
	srv.Router = router.SetupRouter(db.DB)

	log.Fatal(srv.Start(":8080"))
}
