package models

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/account-type-enum"
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
