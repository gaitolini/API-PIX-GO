package services

import (
	"api-pix-go/services/bb"
	"api-pix-go/services/bradesco"
	"fmt"
	"time"
)

type PixService interface {
	CreatePix(requestData map[string]interface{}, txid string) (string, error)
	GetPixStatus(txid string) (string, error)
	GetPixList(startTime, endTime time.Time) (string, error)
	GetPixByTxID(txid string) (string, error) // Novo método
}

func GetPixService(bank string) (PixService, error) {
	switch bank {
	case "bb":
		return bb.NewBBService(), nil
	case "bradesco":
		return bradesco.NewBradescoService(), nil
	default:
		return nil, fmt.Errorf("Banco não suportado: %s", bank)
	}
}
