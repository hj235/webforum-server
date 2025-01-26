package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hj235/cvwo/internal/database"
	"github.com/hj235/cvwo/internal/router"
)

func main() {
	database.InitialiseDB()

	serverPort := os.Getenv("SERVER_PORT")
	r := router.Setup()
	log.Printf("Listening on port %s!\n", serverPort)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", serverPort), r))
}
