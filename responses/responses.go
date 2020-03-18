package responses

import "github.com/SolarLabRU/fastpay-go-commons/models"

type BankResponse struct {
	Data models.Bank
	BaseResponse
}

type AccountResponse struct {
	Data models.Account
	BaseResponse
}

type AccountListResponse struct {
	Data []models.Account
	BaseResponse
}
