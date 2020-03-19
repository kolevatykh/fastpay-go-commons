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
	Data AccountDataPageResponse `json:"data"`
	BaseResponse
}

type AccountDataPageResponse struct {
	Metadata peer.QueryResponseMetadata `json:"metadata"`
	Items    []models.Account           `json:"items"`
}
