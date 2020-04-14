package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
	"github.com/SolarLabRU/fastpay-go-commons/requests"
)

type CrossTransferRequest struct {
	Routes             []models.CurrencyContractRoutingItem `json:"routes" validate:"required,dive"`
	AddressFrom        string                               `json:"addressFrom" validate:"required,validHex40"`
	To                 string                               `json:"to" validate:"required,validHex40or64"`
	Amount             int64                                `json:"amount" validate:"required"`
	CurrencyCodeFrom   int                                  `json:"currencyCodeFrom" validate:"required,min=0"`
	CurrencyCodeTo     int                                  `json:"currencyCodeTo" validate:"required,min=0"`
	CustomerIdentifier string                               `json:"customerIdentifier" validate:"omitempty,validHex64"`
	CountryCode        string                               `json:"countryCode" validate:"omitempty,min=3,max=3"`
	Payload            string                               `json:"payload"`
	MsgHash            string                               `json:"msgHash" validate:"required"`
	Sig                requests.SignDto                     `json:"sig" validate:"required,dive"`
	TransactionId      int                                  `json:"transactionId" validate:"required,min=0"`
}
