package bb

import (
	"context"
	"fmt"

	"api-pix-go/config"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/clientcredentials"
)

type BBService struct{}

func NewBBService() *BBService {
	return &BBService{}
}

func (b *BBService) CreatePix(amount string, txid string, chavePix string) (string, error) {
	// Configuração OAuth2 manual para o Banco do Brasil
	oauthConfig := clientcredentials.Config{
		ClientID:     config.GetBBClientID(),
		ClientSecret: config.GetBBClientSecret(),
		TokenURL:     config.GetBBTokenURL(),
		Scopes: []string{
			"cob.write",
			"cob.read",
			"cobv.write",
			"cobv.read",
			"lotecobv.write",
			"lotecobv.read",
			"pix.write",
			"pix.read",
			"webhook.read",
			"webhook.write",
			"payloadlocation.write",
			"payloadlocation.read",
		},
	}

	// Obter o token OAuth2
	token, err := oauthConfig.Token(context.Background())
	if err != nil {
		return "", fmt.Errorf("Erro ao obter token OAuth2: %v", err)
	}

	// Criar cliente HTTP com Resty e adicionar o token OAuth2 manualmente
	client := resty.New().
		SetAuthToken(token.AccessToken) // Adicionar o token OAuth2 no cabeçalho

	// Corpo da requisição para criar o PIX
	requestBody := map[string]interface{}{
		"calendario": map[string]interface{}{
			"expiracao": 3600,
		},
		"valor": map[string]interface{}{
			"original": amount,
		},
		"chave": chavePix,
	}

	// Fazer a requisição para criar o PIX
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("developer_application_key", config.GetBBApplicationKey()). // Chave do Banco do Brasil
		SetBody(requestBody).
		Post(fmt.Sprintf("%s/cob/%s", config.GetBBApiURL(), txid))

	if err != nil {
		return "", fmt.Errorf("Erro ao criar PIX no Banco do Brasil: %v", err)
	}

	return response.String(), nil
}

func (b *BBService) GetPixStatus(txid string) (string, error) {
	// Configuração OAuth2 manual
	oauthConfig := clientcredentials.Config{
		ClientID:     config.GetBBClientID(),
		ClientSecret: config.GetBBClientSecret(),
		TokenURL:     config.GetBBTokenURL(),
	}

	// Obter o token OAuth2
	token, err := oauthConfig.Token(context.Background())
	if err != nil {
		return "", fmt.Errorf("Erro ao obter token OAuth2: %v", err)
	}

	// Criar cliente HTTP com Resty e adicionar o token OAuth2 manualmente
	client := resty.New().
		SetAuthToken(token.AccessToken) // Adicionar o token OAuth2 no cabeçalho

	// Fazer a requisição para obter o status do PIX
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("developer_application_key", config.GetBBApplicationKey()). // Chave do Banco do Brasil
		Get(fmt.Sprintf("%s/cob/%s", config.GetBBApiURL(), txid))

	if err != nil {
		return "", fmt.Errorf("Erro ao consultar status do PIX no Banco do Brasil: %v", err)
	}

	return response.String(), nil
}
