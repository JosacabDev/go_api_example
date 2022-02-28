package main

import (
	"log"
	"net/http"

	"github.com/JosacabDev/go_api/db"
	"github.com/JosacabDev/go_api/handler"
)

func main() {
	dbMemory := db.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &dbMemory)

	log.Println("Server Up: Port: 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("SERVER INTERNAL ERROR: %v\n", err)
	}
}
