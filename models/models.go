package models

import (
	"sort"

	access_role_enum "github.com/SolarLabRU/fastpay-go-commons/enums/access-role-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/account-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/cross-transaction-payload-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/cross-transaction-status-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contracts-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/deal-state-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/deal-transfer-status-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/invite-status-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/invoice-state-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/member-deal-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/operation-deal-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/safe-deal-deposit-type-enum"
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

type Arbitrator struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	DocType string `json:"docType"`
}

type Bank struct {
	Address   string                        `json:"address"`
	Name      string                        `json:"name"`
	BIK       string                        `json:"bik"`
	State     state_enum.State              `json:"state"`
	CreatedBy string                        `json:"createdBy"`
	Encrypted bool                          `json:"encrypted"`
	MSPId     string                        `json:"mspId"`
	Roles     []access_role_enum.AccessRole `json:"roles"`
	Conf      string                        `json:"conf"`
	DocType   string                        `json:"docType"`
}

type Currency struct {
	Code       int                             `json:"code"`
	Name       string                          `json:"name"`
	Type       currency_type_enum.CurrencyType `json:"type"`
	Unit       string                          `json:"unit"`
	Symbol     string                          `json:"symbol"`
	Decimals   int                             `json:"decimals"`
	Properties map[string]string               `json:"properties"`
	Enabled    bool                            `json:"enabled"`
	DocType    string                          `json:"docType"`
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
	ErrorMessage  string                                    `json:"errorMessage"`
	Payload       string                                    `json:"payload"`
	Timestamp     int64                                     `json:"timestamp"`
	TransactionId string                                    `json:"transactionId"`
	SenderAddress string                                    `json:"senderAddress"`
	Data          string                                    `json:"data"`
	InvoiceNumber string                                    `json:"invoiceNumber"`
	OrdinalNumber int                                       `json:"ordinalNumber"`
}

type TransactionHistoryEvent struct {
	History TransactionHistory `json:"history"`
}

type ExecutedTransactionCurrencyExchangeContractItem struct {
	From         string                                                 `json:"from"`
	To           string                                                 `json:"to"`
	CurrencyCode int                                                    `json:"currencyCode"`
	Amount       int64                                                  `json:"amount"`
	Payload      cross_transaction_payload_enum.CrossTransactionPayload `json:"payload"`
}

type CrossTransactionHistory struct {
	Routes               []CurrencyContractRoutingItem                        `json:"routes"`
	AddressFrom          string                                               `json:"addressFrom"`
	Timestamp            int64                                                `json:"timestamp"`
	TransactionId        string                                               `json:"transactionId"`
	Amount               int64                                                `json:"amount"`
	Payload              string                                               `json:"payload"`
	To                   string                                               `json:"to"`
	EncryptedSecretKeys  []AccountSecretKey                                   `json:"encryptedSecretKeys"`
	CurrencyCodeFrom     int                                                  `json:"currencyCodeFrom"`
	CurrencyCodeTo       int                                                  `json:"currencyCodeTo"`
	CustomerIdentifier   string                                               `json:"customerIdentifier"`
	CountryCode          string                                               `json:"countryCode"`
	Details              []ExecutedTransactionCurrencyExchangeContractItem    `json:"details"`
	SenderAddress        string                                               `json:"senderAddress"`
	BankAddress          string                                               `json:"bankAddress"`
	Status               cross_transaction_status_enum.CrossTransactionStatus `json:"status"`
	TxId                 string                                               `json:"txId"`
	ErrorCode            int                                                  `json:"errorCode"`
	ErrorMessage         string                                               `json:"errorMessage"`
	Data                 string                                               `json:"data"`
	TransactionHistories []TransactionHistory                                 `json:"transactionHistories"`
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
	CurrencyCode    int    `json:"currencyCode" valid:"required,range(0|999)~ErrorCurrencyCodeRange"`
	BankClaims      string `json:"bankClaims" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	BankLiabilities string `json:"bankLiabilities" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
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
	Address         string            `json:"address"`
	BankDisplayName string            `json:"bankDisplayName"`
	State           state_enum.State  `json:"state"`
	CountryCode     string            `json:"countryCode"`
	Owner           string            `json:"owner"`
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
	BankAddress         string `json:"bankAddress"`
	BankDisplayName     string `json:"bankDisplayName"`
	CountryCode         string `json:"countryCode"`
	CustomerDisplayName string `json:"customerDisplayName"`
	DocType             string `json:"docType"`
}

type CurrencyExchangeContractMutable struct {
	Id                   string                                                              `json:"id" valid:"required"`
	AddressAccountSell   string                                                              `json:"addressAccountSell" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	AddressAccountBuy    string                                                              `json:"addressAccountBuy" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	AddressCommission    string                                                              `json:"addressCommission" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyCodeSell     int                                                                 `json:"currencyCodeSell" valid:"range(0|999)~ErrorCurrencyCodeRange"`
	CurrencyCodeBuy      int                                                                 `json:"currencyCodeBuy" valid:"range(0|999)~ErrorCurrencyCodeRange"`
	CurrencySymbolSell   string                                                              `json:"currencySymbolSell"`
	CurrencySymbolBuy    string                                                              `json:"currencySymbolBuy"`
	CurrencyUnitSell     string                                                              `json:"currencyUnitSell"`
	CurrencyUnitBuy      string                                                              `json:"currencyUnitBuy"`
	Type                 currency_exchange_contracts_type_enum.CurrencyExchangeContractsType `json:"type"`
	Price                float64                                                             `json:"price" valid:"range(0|9223372036854775807)"`
	FractionalCommission float64                                                             `json:"fractionalCommission" valid:"range(0|1)"`
	MaxCommission        int64                                                               `json:"maxCommission" valid:"range(0|9223372036854775807)"`
	MinAmount            int64                                                               `json:"minAmount" valid:"range(0|9223372036854775807)"`
	MaxAmount            int64                                                               `json:"maxAmount" valid:"range(0|9223372036854775807)"`
	StartDate            int64                                                               `json:"startDate" valid:"range(0|9223372036854775807)"`
	EndDate              int64                                                               `json:"endDate" valid:"range(0|9223372036854775807)"`
}

type CurrencyExchangeContract struct {
	CurrencyExchangeContractMutable
	BankAddress     string `json:"bankAddress"`
	BankDisplayName string `json:"bankDisplayName"`
	DocType         string `json:"docType"`
}

type BaseEvent struct {
	ChaincodeName string `json:"chaincodeName"`
	FunctionName  string `json:"functionName"`
}

type EventBatch struct {
	Events []EventBatchItem `json:"events"`
}

type EventBatchItem struct {
	EventName string      `json:"eventName"`
	Data      interface{} `json:"data"`
}

type CurrencyEvent struct {
	BaseEvent
	Data Currency `json:"data"`
}

type AccountEvent struct {
	BaseEvent
	Data Account `json:"data"`
}

type ArbitratorEvent struct {
	BaseEvent
	Data Arbitrator `json:"data"`
}

type BankEvent struct {
	BaseEvent
	Data Bank `json:"data"`
}

type AvailablePlatformsEvent struct {
	BaseEvent
	Data string `json:"data"`
}

type TransactionEvent struct {
	BaseEvent
	Data []TransactionHistory `json:"data"`
}

type CrossTransactionEvent struct {
	BaseEvent
	Data CrossTransactionHistory `json:"data"`
}

type ClaimsItemResponse struct {
	CurrencyCode    int      `json:"currencyCode"`
	BankClaims      BankInfo `json:"bankClaims"`
	BankLiabilities BankInfo `json:"bankLiabilities"`
	Amount          int64    `json:"amount"`
	Unconfirmed     int64    `json:"unconfirmed"`
}

type ClearingData struct {
	Id          string               `json:"id"`
	Owner       string               `json:"owner"`
	Claims      int64                `json:"claims"`
	Liabilities int64                `json:"liabilities"`
	History     []ClaimsItemResponse `json:"history"`
	Netting     []AmountOfBank       `json:"netting"`
	Procedure   []ClaimsItemResponse `json:"procedure"`
	Created     int64                `json:"created"`
}

type ClearingEvent struct {
	BaseEvent
	Data ClearingData `json:"data"`
}

type ClaimsEvent struct {
	BaseEvent
	Data []ClaimsItem `json:"data"`
}

type ClientBankEvent struct {
	BaseEvent
	Data ClientBank `json:"data"`
}

type CustomerEvent struct {
	BaseEvent
	Data Customer `json:"data"`
}

type ContractEvent struct {
	BaseEvent
	Data CurrencyExchangeContract `json:"data"`
}

type SetBalanceAccountParam struct {
	Address     string
	AddressBank string
	Value       int64
	Operation   string
	TxId        string
}

type Invoice struct {
	Number       string                          `json:"number"`
	CurrencyCode int                             `json:"currencyCode"`
	Amount       int64                           `json:"amount"`
	Description  string                          `json:"description"`
	Recipient    string                          `json:"recipient"`
	Payer        string                          `json:"payer"`
	State        invoice_state_enum.InvoiceState `json:"state"`
	Created      int64                           `json:"created"`
	Updated      int64                           `json:"updated"`
	ErrorCode    int                             `json:"errorCode"`
	Owner        string                          `json:"owner"`
	DocType      string                          `json:"docType"`
}

type InvoiceEvent struct {
	BaseEvent
	Data Invoice `json:"data"`
}

type LimitsAccount struct {
	Daily   int64 `json:"daily"`
	Monthly int64 `json:"monthly"`
}

type Deal struct {
	Id                    string                            `json:"id"`
	OfferId               string                            `json:"offerId"`
	Owner                 string                            `json:"owner"`
	State                 deal_state_enum.DealState         `json:"state"`
	Terms                 TermsDeal                         `json:"terms"`
	ActualTerms           TermsDeal                         `json:"actualTerms"`
	Invitations           map[string]*Invitation            `json:"invitations"`
	Participants          map[string]*Participant           `json:"participants"`
	TermsContractConclude map[string]*TermsContractConclude `json:"termsContractConclude"`
	NecessaryTransfers    map[string]*TransferSafeDeal      `json:"necessaryTransfers"`
	DepositHistory        []DepositSafeDealHistory          `json:"depositHistory"`
}

type TransferSafeDeal struct {
	AddressFrom  string           `json:"addressFrom"`
	AddressTo    string           `json:"addressTo"`
	CurrencyInfo CurrencyDealInfo `json:"currencyInfo"`
	Amount       int64            `json:"amount" valid:"optional,range(0|9223372036854775807)"`
}

type TermsContractConclude struct {
	AddressFrom      string                               `json:"addressFrom"`
	AddressTo        string                               `json:"addressTo"`
	EscrowAddress    string                               `json:"escrowAddress"`
	MemberTypeTo     member_deal_type_enum.MemberDealType `json:"memberTypeTo"`
	TxId             string                               `json:"txId"`
	IsComplete       bool                                 `json:"isComplete"`
	CurrencyInfo     CurrencyDealInfo                     `json:"currencyInfo"`
	ObligatoryAmount int64                                `json:"obligatoryAmount"`
	CurrentAmount    int64                                `json:"currentAmount"`
	NeedAmount       int64                                `json:"needAmount"`
}

type Invitation struct {
	AddressFrom        string                               `json:"addressFrom"`
	InviteAddress      string                               `json:"inviteAddress"`
	InviteCurrencyInfo CurrencyDealInfo                     `json:"inviteCurrencyInfo"`
	Created            int64                                `json:"created"`
	InviteStatus       invite_status_enum.InviteStatus      `json:"inviteStatus"`
	MemberType         member_deal_type_enum.MemberDealType `json:"memberType"`
}

type Participant struct {
	Address        string                                       `json:"address"`
	CurrencyInfo   CurrencyDealInfo                             `json:"currencyInfo"`
	MemberType     member_deal_type_enum.MemberDealType         `json:"memberType"`
	Created        int64                                        `json:"created"`
	WasInvited     bool                                         `json:"wasInvited"`
	TransferStatus deal_transfer_status_enum.DealTransferStatus `json:"transferStatus"`
}

type TermsDeal struct {
	AddressInitiator       string                                     `json:"addressInitiator" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyInfoInitiator  CurrencyDealInfo                           `json:"currencyInfoInitiator" valid:"required"`
	AmountInitiator        int64                                      `json:"amountInitiator" valid:"range(0|9223372036854775807)"` // Сумму которую инициатор отдает
	OperationTypeInitiator operation_deal_type_enum.OperationDealType `json:"operationTypeInitiator" valid:"required,range(0|2)"`   // В каком виде инициатор отдает указанную сумму
	Price                  float64                                    `json:"price" valid:"required,range(0|9223372036854775807)~ErrorAmountNegative"`
	MinAmount              int64                                      `json:"minAmount" valid:"required,range(0|9223372036854775807)~ErrorAmountNegative"`
	MaxAmount              int64                                      `json:"maxAmount" valid:"required,range(0|9223372036854775807)~ErrorAmountNegative"`
	AddressAcceptor        string                                     `json:"addressAcceptor" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyInfoAcceptor   CurrencyDealInfo                           `json:"currencyInfoAcceptor" valid:"required"`
	OperationTypeAcceptor  operation_deal_type_enum.OperationDealType `json:"operationTypeAcceptor" valid:"required,range(0|2)"`
	AmountAcceptor         int64                                      `json:"amountAcceptor" valid:"optional,range(0|9223372036854775807)~ErrorAmountNegative"`
}

type CurrencyDealInfo struct {
	Code int    `json:"code" valid:"range(0|999)~ErrorCurrencyCodeRange"`
	Name string `json:"name" valid:"required"`
	Unit string `json:"unit"`
}

type SafeDealEvent struct {
	BaseEvent
	Data DealResponseData `json:"data"`
}

type SafeDealDeposit struct {
	SafeDealId         string                   `json:"safeDealId"`
	Deposits           []DepositDetails         `json:"deposits"`
	CurrentBalance     []SetBalanceAccountParam `json:"currentBalance"`
	ForCompleteBalance []SetBalanceAccountParam `json:"forCompleteBalance"`
	AddressTo          string                   `json:"addressTo"`
	CurrencyInfo       CurrencyDealInfo         `json:"currencyInfo"`
	IsComplete         bool                     `json:"isComplete"`
	CurrentAmount      int64                    `json:"currentAmount"`
	NeedAmount         int64                    `json:"needAmount"`
}

type DepositDetails struct {
	AddressFrom string `json:"addressFrom"`
	Amount      int64  `json:"amount"`
	TxID        string `json:"txID"`
}

type DepositSafeDealHistory struct {
	AddressFrom   string                                          `json:"addressFrom"`
	AddressTo     string                                          `json:"addressTo"`
	CurrencyInfo  CurrencyDealInfo                                `json:"currencyInfo"`
	Amount        int64                                           `json:"amount"`
	Type          safe_deal_deposit_type_enum.SafeDealDepositType `json:"type"`
	TxID          string                                          `json:"txID"`
	Timestamp     int64                                           `json:"timestamp"`
	OrdinalNumber int                                             `json:"ordinalNumber"`
}

type DealResponseData struct {
	Id                    string                    `json:"id"`
	OfferId               string                    `json:"offerId"`
	Owner                 string                    `json:"owner"`
	State                 deal_state_enum.DealState `json:"state"`
	Terms                 TermsDeal                 `json:"terms"`
	ActualTerms           TermsDeal                 `json:"actualTerms"`
	Invitations           []Invitation              `json:"invitations"`
	Participants          []Participant             `json:"participants"`
	TermsContractConclude []TermsContractConclude   `json:"termsContractConclude"`
	NecessaryTransfers    []TransferSafeDeal        `json:"necessaryTransfers"`
	DepositHistory        []DepositSafeDealHistory  `json:"depositHistory"`
}

type AccountBalance struct {
	AccountBalanceAddress     string `json:"accountBalanceAddress"`
	AccountBalanceAddressBank string `json:"accountBalanceAddressBank"`
	Op                        string `json:"op"`
	Amount                    int64  `json:"amount"`
	TxID                      string `json:"txID"`
	DocType                   string `json:"docType"`
}

type TransferHistoryLimit struct {
	TransferHistoryAddress string `json:"transferHistoryLimitAddress"`
	YearMonth              int    `json:"yearMonth"`
	Day                    int    `json:"day"`
	Value                  int64  `json:"value"`
	TxID                   string `json:"txID"`
	DocType                string `json:"docType"`
}

type AccountSecretKey struct {
	AddressTo          string `json:"addressTo"`
	EncryptedSecretKey string `json:"encryptedSecretKey"`
}
