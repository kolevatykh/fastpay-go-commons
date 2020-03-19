package responses

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type BankResponse struct {
	Data models.Bank `json:"data"`
	BaseResponse
}

type AccountResponse struct {
	Data models.Account `json:"data"`
	BaseResponse
}

type AccountListResponse struct {
	Data []models.Account `json:"data"`
	BaseResponse
}

type AccountPageResponse struct {
	Data struct {
		Metadata peer.QueryResponseMetadata `json:"metadata"`
		Items    []models.Account           `json:"items"`
	}
	BaseResponse
}

type SenderAddressResponse struct {
	Data struct {
		Address string `json:"address"`
	} `json:"data"`
	BaseResponse
}
