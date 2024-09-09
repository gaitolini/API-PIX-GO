package bradesco

import (
	"api-pix-go/config"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/clientcredentials"
)

type BradescoService struct{}

func NewBradescoService() *BradescoService {
	return &BradescoService{}
}

func (b *BradescoService) getToken() (string, error) {
	oauthConfig := clientcredentials.Config{
		ClientID:     config.GetBradescoClientID(),
		ClientSecret: config.GetBradescoClientSecret(),
		TokenURL:     config.GetBradescoTokenURL(),
		Scopes:       []string{"cob.write", "pix.write"}, // Adicione os escopos necessários
	}

	log.Println("Iniciando a obtenção do token OAuth2 para Bradesco...")

	token, err := oauthConfig.Token(context.Background())
	if err != nil {
		log.Printf("Erro ao obter token OAuth2: %v", err)
		return "", fmt.Errorf("Erro ao obter token OAuth2: %v", err)
	}

	log.Printf("Token obtido com sucesso: %s", token.AccessToken)
	return token.AccessToken, nil
}

func (b *BradescoService) CreatePix(requestData map[string]interface{}, txid string) (string, error) {
	// Obter o token OAuth2
	token, err := b.getToken()
	if err != nil {
		return "", err
	}

	// Criar cliente HTTP com Resty e adicionar o token OAuth2 manualmente
	client := resty.New().SetAuthToken(token)

	// Fazer a requisição para criar o PIX com o JSON recebido
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestData).
		Post(fmt.Sprintf("%s/cob/%s", config.GetBradescoAPIURL(), txid))

	if err != nil {
		return "", fmt.Errorf("Erro ao criar PIX no Bradesco: %v", err)
	}

	return response.String(), nil
}

func (b *BradescoService) GetPixStatus(txid string) (string, error) {
	// Obter o token OAuth2
	log.Println("Tentando consultar o status do PIX no Bradesco...")
	token, err := b.getToken()
	if err != nil {
		return "", err
	}

	// Criar cliente HTTP com Resty e adicionar o token OAuth2 manualmente
	client := resty.New().SetAuthToken(token)

	log.Printf("Consultando o status do PIX com txid: %s", txid)

	// Fazer a requisição para obter o status do PIX
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get(fmt.Sprintf("%s/cob/%s", config.GetBradescoAPIURL(), txid))

	if err != nil {
		log.Printf("Erro ao consultar status do PIX no Bradesco: %v", err)
		return "", fmt.Errorf("Erro ao consultar status do PIX no Bradesco: %v", err)
	}

	log.Printf("Status do PIX consultado com sucesso no Bradesco, resposta: %s", response.String())
	return response.String(), nil
}

func (b *BradescoService) GetPixList(startTime, endTime time.Time) (string, error) {
	// Obter o token OAuth2
	token, err := b.getToken()
	if err != nil {
		return "", err
	}

	// Formatar as datas de início e fim para a URL
	startTimeFormatted := startTime.Format(time.RFC3339)
	endTimeFormatted := endTime.Format(time.RFC3339)

	// Criar cliente HTTP com Resty e adicionar o token OAuth2
	client := resty.New().SetAuthToken(token)

	// Log da URL de requisição
	log.Printf("Consultando lista de cobranças Bradesco de %s até %s", startTimeFormatted, endTimeFormatted)

	// Fazer a requisição para consultar lista de cobranças
	response, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Get(fmt.Sprintf("%s/pix/v2/cob?inicio=%s&fim=%s&gw-dev-app-key=%s",
			config.GetBradescoAPIURL(), startTimeFormatted, endTimeFormatted, config.GetBradescoApplicationKey()))

	if err != nil {
		log.Printf("Erro ao consultar lista de cobranças no Bradesco: %v", err)
		return "", fmt.Errorf("Erro ao consultar lista de cobranças no Bradesco: %v", err)
	}

	log.Printf("Resposta da consulta de lista de cobranças Bradesco: %s", response.String())
	return response.String(), nil
}

func (b *BradescoService) GetPixByTxID(txid string) (string, error) {
	// Obter o token OAuth2
	token, err := b.getToken()
	if err != nil {
		return "", err
	}

	// Criar cliente HTTP com Resty e adicionar o token OAuth2
	client := resty.New().SetAuthToken(token)

	// Fazer a requisição para consultar o PIX pelo txid
	response, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Get(fmt.Sprintf("%s/pix/v2/cob/%s?gw-dev-app-key=%s",
			config.GetBradescoAPIURL(), txid, config.GetBradescoApplicationKey()))

	if err != nil {
		log.Printf("Erro ao consultar PIX no Bradesco: %v", err)
		return "", fmt.Errorf("Erro ao consultar PIX no Bradesco: %v", err)
	}

	log.Printf("Consulta do PIX por txid no Bradesco realizada com sucesso, resposta: %s", response.String())
	return response.String(), nil
}
