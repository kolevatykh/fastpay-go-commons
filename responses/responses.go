package responses

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
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

type Metadata struct {
	FetchedRecordsCount int32  `json:"fetchedRecordsCount"`
	Bookmark            string `json:"bookmark"`
}

type AccountPageData struct {
	Metadata Metadata         `json:"metadata"`
	Items    []models.Account `json:"items"`
}

type AccountPageResponse struct {
	Data AccountPageData `json:"data"`
	BaseResponse
}

type SenderAddressData struct {
	Address string `json:"address"`
}

type SenderAddressResponse struct {
	Data SenderAddressData `json:"data"`
	BaseResponse
}
