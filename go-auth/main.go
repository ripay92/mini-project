package main

import (
	models "GO-AUTH/model"
	"GO-AUTH/router"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	models.ConnectDatabase()
	r:= router.NewRouter()

	log.Println("Server started at :9000")
	err := http.ListenAndServe(":9000", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(r))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}