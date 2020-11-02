package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

type CrossTransferRequest struct {
	Routes              []models.CurrencyContractRoutingItem `json:"routes" valid:"required"`
	AddressFrom         string                               `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	To                  string                               `json:"to" valid:"validHex40or64~ErrorAddressNotFollowingRegex"`
	EncryptedSecretKeys []models.AccountSecretKey            `json:"encryptedSecretKeys"`
	Amount              int64                                `json:"amount" valid:"required~ErrorAmountNotPassed"`
	CurrencyCodeFrom    int                                  `json:"currencyCodeFrom" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	CurrencyCodeTo      int                                  `json:"currencyCodeTo" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	CustomerIdentifier  string                               `json:"customerIdentifier" valid:"validHex64~ErrorIdentifierNotFolowingRegex"`
	CountryCode         string                               `json:"countryCode" valid:"stringlength(3|3)"`
	Payload             string                               `json:"payload"`
	MsgHash             string                               `json:"msgHash" valid:"required"`
	Sig                 SignDto                              `json:"sig" valid:"required"`
	Exp                 int64                                `json:"exp" valid:"required~ErrorTimestampNotPassed"`
	TransactionId       string                               `json:"transactionId" valid:"required~ErrorTransactionIdNotPassed,uuidv4"`
	BankAddress         string                               `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}

func (crossTransferRequest *CrossTransferRequest) SetDefaults() {
	if crossTransferRequest.EncryptedSecretKeys == nil {
		crossTransferRequest.EncryptedSecretKeys = []models.AccountSecretKey{}
	}
}
