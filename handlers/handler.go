package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Artic-Dev/ArticStyleApi-GO/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Controller() {
	router := mux.NewRouter()

	router.HandleFunc("/", routers.Index)
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handlers := cors.AllowAll().Handler(router)
	log.Println("Servidor funcionando por el puerto :"+PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handlers))
}