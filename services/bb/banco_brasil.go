package bb

import (
	"context"
	"fmt"
	"log" // Adicionado para os logs
	"time"

	"api-pix-go/config"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/clientcredentials"
)

type BBService struct{}

func NewBBService() *BBService {
	return &BBService{}
}

// Função para obter o token OAuth2 com logs
func (b *BBService) getToken() (string, error) {
	oauthConfig := clientcredentials.Config{
		ClientID:     config.GetBBClientID(),
		ClientSecret: config.GetBBClientSecret(),
		TokenURL:     config.GetBBTokenURL(),
		Scopes: []string{
			"cob.write", "cob.read", "cobv.write", "cobv.read",
			"lotecobv.write", "lotecobv.read", "pix.write", "pix.read",
			"webhook.read", "webhook.write", "payloadlocation.write", "payloadlocation.read",
		},
	}

	log.Println("Iniciando a obtenção do token OAuth2...")

	token, err := oauthConfig.Token(context.Background())
	if err != nil {
		log.Printf("Erro ao obter token OAuth2: %v", err)
		return "", fmt.Errorf("Erro ao obter token OAuth2: %v", err)
	}

	log.Printf("Token obtido com sucesso: %s", token.AccessToken)
	return token.AccessToken, nil
}

// Função para criar o PIX com logs
func (b *BBService) CreatePix(requestData map[string]interface{}, txid string) (string, error) {
	// Obter o token OAuth2
	token, err := b.getToken()
	if err != nil {
		return "", err
	}

	// Log do corpo da requisição antes de enviar
	log.Printf("Enviando requisição para criar PIX no BB com o seguinte body: %+v", requestData)

	// Criar cliente HTTP com Resty e adicionar o token OAuth2
	client := resty.New().SetAuthToken(token)

	// Fazer a requisição para criar o PIX com o JSON recebido
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestData).
		Post(fmt.Sprintf("%s/pix/v2/cob?gw-dev-app-key=%s", config.GetBBApiURL(), config.GetBBApplicationKey())) // Passa "gw-dev-app-key" na query

	if err != nil {
		log.Printf("Erro ao criar PIX no Banco do Brasil: %v", err)
		return "", fmt.Errorf("Erro ao criar PIX no Banco do Brasil: %v", err)
	}

	log.Printf("Resposta da criação do PIX: %s", response.String())
	return response.String(), nil
}

// Função para consultar o status do PIX com logs
func (b *BBService) GetPixStatus(txid string) (string, error) {
	// Obter o token OAuth2
	log.Println("Tentando consultar o status do PIX...")
	token, err := b.getToken()
	if err != nil {
		return "", err
	}

	// Criar cliente HTTP com Resty e adicionar o token OAuth2
	client := resty.New().SetAuthToken(token)

	log.Printf("Consultando o status do PIX com txid: %s", txid)

	// Fazer a requisição para obter o status do PIX
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get(fmt.Sprintf("%s/pix/v2/cob/%s?gw-dev-app-key=%s", config.GetBBApiURL(), txid, config.GetBBApplicationKey())) // Passa "gw-dev-app-key" na query

	if err != nil {
		log.Printf("Erro ao consultar status do PIX no Banco do Brasil: %v", err)
		return "", fmt.Errorf("Erro ao consultar status do PIX no Banco do Brasil: %v", err)
	}

	log.Printf("Status do PIX consultado com sucesso, resposta: %s", response.String())
	return response.String(), nil
}

// Função para consultar lista de cobranças imediatas
func (b *BBService) GetPixList(startTime, endTime time.Time) (string, error) {
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
	log.Printf("Consultando lista de cobranças de %s até %s", startTimeFormatted, endTimeFormatted)

	// Fazer a requisição para consultar lista de cobranças
	response, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Get(fmt.Sprintf("%s/pix/v2/cob?inicio=%s&fim=%s&gw-dev-app-key=%s",
			config.GetBBApiURL(), startTimeFormatted, endTimeFormatted, config.GetBBApplicationKey()))

	if err != nil {
		log.Printf("Erro ao consultar lista de cobranças: %v", err)
		return "", fmt.Errorf("Erro ao consultar lista de cobranças: %v", err)
	}

	log.Printf("Resposta da consulta de lista de cobranças: %s", response.String())
	return response.String(), nil
}

func (b *BBService) GetPixByTxID(txid string) (string, error) {
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
			config.GetBBApiURL(), txid, config.GetBBApplicationKey()))

	if err != nil {
		log.Printf("Erro ao consultar PIX no Banco do Brasil: %v", err)
		return "", fmt.Errorf("Erro ao consultar PIX no Banco do Brasil: %v", err)
	}

	log.Printf("Consulta do PIX por txid realizada com sucesso, resposta: %s", response.String())
	return response.String(), nil
}
