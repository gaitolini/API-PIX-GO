package bradesco

import (
	"api-pix-go/config"
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/clientcredentials"
)

type BradescoService struct{}

func NewBradescoService() *BradescoService {
	return &BradescoService{}
}

func (b *BradescoService) CreatePix(amount string, txid string, chavePix string) (string, error) {
	// Configuração OAuth2 manual
	oauthConfig := clientcredentials.Config{
		ClientID:     config.GetBradescoClientID(),
		ClientSecret: config.GetBradescoClientSecret(),
		TokenURL:     config.GetBradescoTokenURL(),
		Scopes:       []string{}, // Adicione escopos necessários se houver
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
		SetBody(requestBody).
		Post(fmt.Sprintf("%s/cob/%s", config.GetBradescoAPIURL(), txid))

	if err != nil {
		return "", fmt.Errorf("Erro ao criar PIX no Bradesco: %v", err)
	}

	return response.String(), nil
}

func (b *BradescoService) GetPixStatus(txid string) (string, error) {
	// Configuração OAuth2 manual
	oauthConfig := clientcredentials.Config{
		ClientID:     config.GetBradescoClientID(),
		ClientSecret: config.GetBradescoClientSecret(),
		TokenURL:     config.GetBradescoTokenURL(),
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
		Get(fmt.Sprintf("%s/cob/%s", config.GetBradescoAPIURL(), txid))

	if err != nil {
		return "", fmt.Errorf("Erro ao consultar status do PIX no Bradesco: %v", err)
	}

	return response.String(), nil
}
