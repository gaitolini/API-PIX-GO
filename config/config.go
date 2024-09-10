package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Função para carregar o arquivo .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
}

// Função para obter o client_id do Banco do Brasil
func GetBBClientID() string {
	return os.Getenv("BB_CLIENT_ID")
}

// Função para obter o client_secret do Banco do Brasil
func GetBBClientSecret() string {
	return os.Getenv("BB_CLIENT_SECRET")
}

// Função para obter o developer_application_key do Banco do Brasil
func GetBBApplicationKey() string {
	return os.Getenv("BB_APPLICATION_KEY")
}

// Função para obter a URL do token OAuth2 do Banco do Brasil
func GetBBTokenURL() string {
	return os.Getenv("BB_TOKEN_URL")
}

// Função para obter a URL base da API do Banco do Brasil
func GetBBApiURL() string {
	return os.Getenv("BB_API_URL")
}

// Função para obter o developer_application_key do Banco do Brasil
func GetBradescoApplicationKey() string {
	return os.Getenv("BRADESCO_APPLICATION_KEY")
}

// Função para obter o client_id do Bradesco
func GetBradescoClientID() string {
	return os.Getenv("BRADESCO_CLIENT_ID")
}

// Função para obter o client_secret do Bradesco
func GetBradescoClientSecret() string {
	return os.Getenv("BRADESCO_CLIENT_SECRET")
}

// Função para obter a URL do token OAuth2 do Bradesco
func GetBradescoTokenURL() string {
	return os.Getenv("BRADESCO_TOKEN_URL")
}

// Função para obter a URL base da API do Bradesco
func GetBradescoAPIURL() string {
	return os.Getenv("BRADESCO_API_URL")
}

func GetPort() string {
	return os.Getenv("PORT")
}
