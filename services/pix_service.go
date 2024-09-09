package services

import (
	"api-pix-go/services/bb"
	"api-pix-go/services/bradesco"
	"fmt"
)

type PixService interface {
	CreatePix(amount string, txid string, chavePix string) (string, error)
	GetPixStatus(txid string) (string, error)
}

func GetPixService(bank string) (PixService, error) {
	switch bank {
	case "bradesco":
		return bradesco.NewBradescoService(), nil
	case "bb":
		return bb.NewBBService(), nil
	default:
		return nil, fmt.Errorf("Banco não suportado: %s", bank)
	}
}
