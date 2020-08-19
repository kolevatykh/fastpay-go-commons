package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

type CrossTransferRequest struct {
	Routes             []models.CurrencyContractRoutingItem `json:"routes" valid:"required"`
	AddressFrom        string                               `json:"addressFrom" valid:"required~60302,validHex40~60301"`
	To                 string                               `json:"to" valid:"validHex40or64~60301"`
	Amount             int64                                `json:"amount" valid:"required~60313"`
	CurrencyCodeFrom   int                                  `json:"currencyCodeFrom" valid:"required~60318,range(0|999)~60317"`
	CurrencyCodeTo     int                                  `json:"currencyCodeTo" valid:"required~60318,range(0|999)~60317"`
	CustomerIdentifier string                               `json:"customerIdentifier" valid:"validHex64~60303"`
	CountryCode        string                               `json:"countryCode" valid:"stringlength(3|3)"`
	Payload            string                               `json:"payload"`
	MsgHash            string                               `json:"msgHash" valid:"required"`
	Sig                SignDto                              `json:"sig" valid:"required"`
	Exp                int64                                `json:"exp" valid:"required~60332"`
	TransactionId      string                               `json:"transactionId" valid:"required~60331,uuidv4"`
}
