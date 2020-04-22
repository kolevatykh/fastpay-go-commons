package models

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/account-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/cross-transaction-payload-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/cross-transaction-status-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/transaction-status-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/transaction-type-enum"
	"sort"
)

type Account struct {
	Address                    string                            `json:"address"`
	State                      state_enum.State                  `json:"state"`
	CurrencyCode               int                               `json:"currencyCode"`
	JuridicalType              juridical_type_enum.JuridicalType `json:"juridicalType"`
	BikBankSetterJuridicalType string                            `json:"bikBankSetterJuridicalType"`
	IdentityType               identity_type_enum.IdentityType   `json:"identityType"`
	Owner                      string                            `json:"owner"`
	Type                       account_type_enum.AccountType     `json:"type"`
	Identifiers                []string                          `json:"identifiers"`
	Encrypted                  bool                              `json:"encrypted"`
	Created                    int64                             `json:"created"`
	PublicKey                  string                            `json:"publicKey"`
	DocType                    string                            `json:"docType"`
}

type Bank struct {
	Address     string           `json:"address"`
	Name        string           `json:"name"`
	BIK         string           `json:"bik"`
	State       state_enum.State `json:"state"`
	CreatedBy   string           `json:"createdBy"`
	IsOwner     bool             `json:"isOwner"`
	Encrypted   bool             `json:"encrypted"`
	IsRegulator bool             `json:"isRegulator"`
	MSPId       string           `json:"mspId"`
	Conf        string           `json:"conf"`
	DocType     string           `json:"docType"`
}

type Currency struct {
	Code     int    `json:"code"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Enabled  bool   `json:"enabled"`
	DocType  string `json:"docType"`
}

type AmountOfBank struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}

type TransactionHistory struct {
	TxId          string                                    `json:"txId"`
	AddressFrom   string                                    `json:"addressFrom"`
	AddressTo     string                                    `json:"addressTo"`
	TxType        transaction_type_enum.TransactionType     `json:"txType"`
	Status        transaction_status_enum.TransactionStatus `json:"status"`
	Amount        int64                                     `json:"amount"`
	CurrencyCode  int                                       `json:"currencyCode"`
	ErrorCode     int                                       `json:"errorCode"`
	Payload       string                                    `json:"payload"`
	Timestamp     int64                                     `json:"timestamp"`
	TransactionId int                                       `json:"transactionId"`
	SenderAddress string                                    `json:"senderAddress"`
	Data          string                                    `json:"data"`
}

type TransactionHistoryEvent struct {
	History TransactionHistory `json:"history"`
}

type ExecutedTransactionCurrencyExchangeContractItem struct {
	From         string                                                 `json:"from"`
	To           string                                                 `json:"to"`
	CurrencyCode int                                                    `json:"currencyCode"`
	Amount       int64                                                  `json:"amount"`
	Payload      cross_transaction_payload_enum.CrossTransactionPayload `json:"payload"` // TODO add enum CrossTransactionPayload
}

type CrossTransactionHistory struct {
	Routes               []CurrencyContractRoutingItem                        `json:"routes"`
	AddressFrom          string                                               `json:"addressFrom"`
	Timestamp            int64                                                `json:"timestamp"`
	TransactionId        int                                                  `json:"transactionId"`
	Amount               int64                                                `json:"amount"`
	Payload              string                                               `json:"payload"`
	To                   string                                               `json:"to"`
	CurrencyCodeFrom     int                                                  `json:"currencyCodeFrom"`
	CurrencyCodeTo       int                                                  `json:"currencyCodeTo"`
	CustomerIdentifier   string                                               `json:"customerIdentifier"`
	CountryCode          string                                               `json:"countryCode"`
	Details              []ExecutedTransactionCurrencyExchangeContractItem    `json:"details"`
	SenderAddress        string                                               `json:"senderAddress"`
	BankId               string                                               `json:"bankId"`
	Status               cross_transaction_status_enum.CrossTransactionStatus `json:"status"`
	TxId                 string                                               `json:"txId"`
	ErrorCode            int                                                  `json:"errorCode"`
	Data                 string                                               `json:"data"`
	TransactionHistories []TransactionHistory
}

type CrossTransactionHistoryEvent struct {
	Data CrossTransactionHistory `json:"history"`
}

type CurrencyContractRoutingItem struct {
	CurrencyExchangeContract
	AmountInput  int64 `json:"amountInput"`
	AmountOutput int64 `json:"amountOutput"`
}

type DepositedBalance struct {
	WithdrawResult
	IssueBankAddress string           `json:"issueBankAddress"`
	BanksBalance     map[string]int64 `json:"banksBalance"`
}

type WithdrawResult struct {
	AccountAmount    int64  `json:"accountAmount"`
	IssueBankAmount  int64  `json:"issueBankAmount"`
	BanksTotalAmount int64  `json:"banksTotalAmount"`
	TxId             string `json:"txId"`
}

type ClaimsItem struct {
	CurrencyCode    int    `json:"currencyCode" validate:"required,gte=0"`
	BankClaims      string `json:"bankClaims" validate:"required,validHex40"`
	BankLiabilities string `json:"bankLiabilities" validate:"required,validHex40"`
	Amount          int64  `json:"amount"`
	Unconfirmed     int64  `json:"unconfirmed"`
}

type ClaimsItemDoc struct {
	ClaimsItem
	DocType string `json:"docType"`
	Id      string `json:"id"`
}

type BankInfo struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Bik     string `json:"bik"`
}

type ClearingInfo struct {
	Id          string                   `json:"id"`
	Banks       map[string]*ClearingBank `json:"banks"`
	Owner       string                   `json:"owner"`
	Claims      int64                    `json:"claims"`
	Liabilities int64                    `json:"liabilities"`
	History     []ClaimsItem             `json:"history"`
	Netting     map[string]int64         `json:"netting"`
	Procedure   []ClaimsItem             `json:"procedure"`
	Created     int64                    `json:"created"`
	DocType     string                   `json:"docType"`
}

func (ci *ClearingInfo) GetSortBanksKeys() []string {
	keys := make([]string, 0)
	for k, _ := range ci.Banks {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func (ci *ClearingInfo) GetSortNettingKeys() []string {
	keys := make([]string, 0)
	for k, _ := range ci.Netting {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

type ClearingBank struct {
	Claims      int64 `json:"claims"`
	Liabilities int64 `json:"liabilities"`
}

type ClaimsAggregate struct {
	Amount      int64    `json:"amount"`
	Unconfirmed int64    `json:"unconfirmed"`
	Ids         []string `json:"ids"`
}

type ClientBank struct {
	BankId          string            `json:"bankId"`
	BankDisplayName string            `json:"bankDisplayName"`
	State           state_enum.State  `json:"state"`
	CountryCode     string            `json:"countryCode"`
	Params          map[string]string `json:"params"`
	DocType         string            `json:"docType"`
}

func (cb *ClientBank) GetSortParamsKeys() []string {
	keys := make([]string, 0)
	for k, _ := range cb.Params {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

type Customer struct {
	Identifier          string `json:"identifier"`
	BankId              string `json:"bankId"`
	BankDisplayName     string `json:"bankDisplayName"`
	CountryCode         string `json:"countryCode"`
	CustomerDisplayName string `json:"customerDisplayName"`
	DocType             string `json:"docType"`
}

type CurrencyExchangeContractMutable struct {
	Id                   string  `json:"id" validate:"required"`
	AddressAccountSell   string  `json:"addressAccountSell" validate:"omitempty,validHex40"`
	AddressAccountBuy    string  `json:"addressAccountBuy" validate:"omitempty,validHex40"`
	AddressCommission    string  `json:"addressCommission" validate:"omitempty,validHex40"`
	CurrencyCodeSell     int     `json:"currencyCodeSell" validate:"omitempty,min=0"`
	CurrencyCodeBuy      int     `json:"currencyCodeBuy" validate:"omitempty,min=0"`
	CurrencySymbolSell   string  `json:"currencySymbolSell" validate:"omitempty,min=3,max=3"`
	CurrencySymbolBuy    string  `json:"currencySymbolBuy" validate:"omitempty,min=3,max=3"`
	Price                float64 `json:"price" validate:"omitempty,gte=0.0000000001"`
	FractionalCommission float64 `json:"fractionalCommission" validate:"omitempty,gte=0,lte=1"`
	MaxCommission        int64   `json:"maxCommission" validate:"omitempty,min=0"`
	MinAmount            int64   `json:"minAmount" validate:"omitempty,min=0"`
	MaxAmount            int64   `json:"maxAmount" validate:"omitempty,min=0"`
	StartDate            int64   `json:"startDate" validate:"omitempty,min=0"`
	EndDate              int64   `json:"endDate" validate:"omitempty,min=0"`
}

type CurrencyExchangeContract struct {
	CurrencyExchangeContractMutable
	BankId          string `json:"bankId"`
	BankDisplayName string `json:"bankDisplayName"`
	DocType         string `json:"docType"`
}

type BaseEvent struct {
	ChaincodeName string `json:"chaincodeName"`
	FunctionName  string `json:"functionName"`
}

type CurrencyEvent struct {
	BaseEvent
	Data models.Currency `json:"data"`
}

type AccountEvent struct {
	BaseEvent
	Data models.Account `json:"data"`
}

type BankEvent struct {
	BaseEvent
	Data models.Bank `json:"data"`
}

type AvailablePlatformsEvent struct {
	BaseEvent
	Data string `json:"data"`
}
