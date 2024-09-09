package controllers

import (
	"api-pix-go/services"
	"api-pix-go/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePix(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]string
	json.NewDecoder(r.Body).Decode(&requestData)

	bank := requestData["bank"]
	amount := requestData["amount"]
	chavePix := requestData["chave_pix"]

	// Gera um txid único
	txid := utils.GenerateTxID()

	// Chamada para GetPixService, lidando com os dois valores retornados
	pixService, err := services.GetPixService(bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Retorna erro se o banco não for suportado
		return
	}

	// Cria o PIX
	result, err := pixService.CreatePix(amount, txid, chavePix)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status":   "success",
		"message":  "PIX created successfully",
		"location": result,
		"txid":     txid,
	})
}

func GetPixStatus(w http.ResponseWriter, r *http.Request) {
	txid := mux.Vars(r)["txid"]
	bank := r.URL.Query().Get("bank")

	// Chamada para GetPixService, lidando com os dois valores retornados
	pixService, err := services.GetPixService(bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Retorna erro se o banco não for suportado
		return
	}

	// Consulta o status do PIX
	result, err := pixService.GetPixStatus(txid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
