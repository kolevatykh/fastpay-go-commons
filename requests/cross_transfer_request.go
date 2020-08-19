package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

type CrossTransferRequest struct {
	Routes             []models.CurrencyContractRoutingItem `json:"routes" valid:"required"`
	AddressFrom        string                               `json:"addressFrom" valid:"required,validHex40"`
	To                 string                               `json:"to" valid:"validHex40or64"`
	Amount             int64                                `json:"amount" valid:"required"`
	CurrencyCodeFrom   int                                  `json:"currencyCodeFrom" valid:"required,min(0)"`
	CurrencyCodeTo     int                                  `json:"currencyCodeTo" valid:"required,min(0)"`
	CustomerIdentifier string                               `json:"customerIdentifier" valid:"validHex64"`
	CountryCode        string                               `json:"countryCode" valid:"min(3),max(3)"`
	Payload            string                               `json:"payload"`
	MsgHash            string                               `json:"msgHash" valid:"required"`
	Sig                SignDto                              `json:"sig" valid:"required"`
	Exp                int64                                `json:"exp" valid:"required"`
	TransactionId      string                               `json:"transactionId" valid:"required,uuidv4"`
}
