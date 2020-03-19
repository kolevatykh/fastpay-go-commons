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

// TODO Если вынести внутрению струкутра Data будет ошибка валидации Cannot use metadata. Metadata did not match schema:
// 1. components.schemas..required: Array must have at least 1 items [recovered]
// При создании чейнкода аккаунтов в тесте
type AccountPageResponse struct {
	Data struct {
		Metadata peer.QueryResponseMetadata `json:"metadata"`
		Items    []models.Account           `json:"items"`
	} `json:"data"`
	BaseResponse
}

type SenderAddressData struct {
	Address string `json:"address"`
}

type SenderAddressResponse struct {
	Data SenderAddressData `json:"data"`
	BaseResponse
}
