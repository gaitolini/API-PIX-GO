package main

import (
	"api-pix-go/config"
	"api-pix-go/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv() // Carregar as variáveis de ambiente

	r := mux.NewRouter()

	r.HandleFunc("/api/pix/create", controllers.CreatePix).Methods("POST")
	r.HandleFunc("/api/pix/status/{txid}", controllers.GetPixStatus).Methods("GET")
	r.HandleFunc("/api/pix/{txid}", controllers.GetPixByTxID).Methods("GET")
	r.HandleFunc("/api/pix/list", controllers.GetPixList).Methods("GET")

	port := config.GetPort()
	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
