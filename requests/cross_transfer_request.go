package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

type CrossTransferRequest struct {
	Routes             []models.CurrencyContractRoutingItem `json:"routes" validate:"required,dive"`
	AddressFrom        string                               `json:"addressFrom" validate:"required,validHex40"`
	To                 string                               `json:"to" validate:"omitempty,validHex40or64"`
	Amount             int64                                `json:"amount" validate:"required"`
	CurrencyCodeFrom   int                                  `json:"currencyCodeFrom" validate:"required,min=0"`
	CurrencyCodeTo     int                                  `json:"currencyCodeTo" validate:"required,min=0"`
	CustomerIdentifier string                               `json:"customerIdentifier" validate:"omitempty,validHex64"`
	CountryCode        string                               `json:"countryCode" validate:"omitempty,min=3,max=3"`
	Payload            string                               `json:"payload"`
	MsgHash            string                               `json:"msgHash" validate:"required"`
	Sig                SignDto                              `json:"sig" validate:"required,dive"`
	Exp                int64                                `json:"exp" validate:"required"`
	TransactionId      int                                  `json:"transactionId" validate:"required,min=0"`
}
