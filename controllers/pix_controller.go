package controllers

import (
	"api-pix-go/services"
	"api-pix-go/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CreatePix(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verifica o banco e remove o campo "bank" do requestData
	bank, ok := requestData["bank"].(string)
	if !ok {
		http.Error(w, "Banco não informado corretamente", http.StatusBadRequest)
		return
	}
	delete(requestData, "bank") // Remove o campo "bank" para não enviá-lo ao banco

	// Gera um txid único
	txid := utils.GenerateTxID()

	// Chama o serviço apropriado com base no banco
	pixService, err := services.GetPixService(bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Cria o PIX
	result, err := pixService.CreatePix(requestData, txid)
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

func GetPixList(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verifica o banco e remove o campo "bank" do requestData
	bank, ok := requestData["bank"].(string)
	if !ok {
		http.Error(w, "Banco não informado corretamente", http.StatusBadRequest)
		return
	}
	delete(requestData, "bank") // Remove o campo "bank" para não enviá-lo ao banco

	// Verifica e converte os parâmetros "inicio" e "fim" do requestData
	startTimeStr, ok := requestData["inicio"].(string)
	if !ok {
		http.Error(w, "Parâmetro 'inicio' inválido ou ausente", http.StatusBadRequest)
		return
	}
	endTimeStr, ok := requestData["fim"].(string)
	if !ok {
		http.Error(w, "Parâmetro 'fim' inválido ou ausente", http.StatusBadRequest)
		return
	}

	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		http.Error(w, "Data de início inválida", http.StatusBadRequest)
		return
	}
	endTime, err := time.Parse(time.RFC3339, endTimeStr)
	if err != nil {
		http.Error(w, "Data de fim inválida", http.StatusBadRequest)
		return
	}

	// Chama o serviço apropriado com base no banco
	pixService, err := services.GetPixService(bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obter a lista de cobranças
	result, err := pixService.GetPixList(startTime, endTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Consulta realizada com sucesso",
		"result":  result,
	})
}

func GetPixByTxID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txid := vars["txid"]

	var requestData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verifica o banco
	bank, ok := requestData["bank"].(string)
	if !ok {
		http.Error(w, "Banco não informado corretamente", http.StatusBadRequest)
		return
	}

	// Chama o serviço apropriado com base no banco
	pixService, err := services.GetPixService(bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Consulta o PIX pelo txid
	result, err := pixService.GetPixByTxID(txid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Consulta realizada com sucesso",
		"result":  result,
	})
}
