package models

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/account-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/transaction-status-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/transaction-type-enum"
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
