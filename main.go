package main

import (
	"log"

	"github.com/SaiKumarMalve/Gator-News/bd"
	"github.com/SaiKumarMalve/Gator-News/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("No connection to the DB")
		return
	}
	handlers.Handlers()
}
