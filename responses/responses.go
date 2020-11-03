package responses

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
	transaction_type_enum "github.com/SolarLabRU/fastpay-go-commons/enums/transaction-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

type BankResponse struct {
	Data models.Bank `json:"data"`
	BaseResponse
}

type MakeSafeDealDepositResponse struct {
	Data                     []models.EventBatchItem `json:"data"`
	NeedMakeTransferRequests []TransferRequest       `json:"transferRequests"`
	BaseResponse
}

type TransferRequest struct {
	AddressFrom  string                                `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	To           string                                `json:"to" valid:"required~ErrorAddressNotPassed,validHex40or64~ErrorAddressOrIdentifierNotFolowingRegex"`
	CurrencyCode int                                   `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Amount       int64                                 `json:"amount" valid:"required~ErrorAmountNotPassed"`
	Payload      string                                `json:"payload"`
	MsgHash      string                                `json:"msgHash"`
	TxType       transaction_type_enum.TransactionType `json:"txType"`
}

type AccountResponse struct {
	Data models.Account `json:"data"`
	BaseResponse
}

type AccountListResponse struct {
	Data []models.Account `json:"data"`
	BaseResponse
}

type ArbitratorListResponse struct {
	Data []models.Arbitrator `json:"data"`
	BaseResponse
}

type Metadata struct {
	FetchedRecordsCount int32  `json:"fetchedRecordsCount"`
	Bookmark            string `json:"bookmark"`
}

type AccountPageData struct {
	Metadata Metadata         `json:"metadata"`
	Items    []models.Account `json:"items"`
}

type AccountPageResponse struct {
	Data AccountPageData `json:"data"`
	BaseResponse
}

type SenderAddressData struct {
	Address string `json:"address"`
}

type SenderAddressResponse struct {
	Data SenderAddressData `json:"data"`
	BaseResponse
}

type BankListResponse struct {
	Data []models.Bank `json:"data"`
	BaseResponse
}

type BankTotalData struct {
	Total int `json:"total"`
}

type BankTotalResponse struct {
	Data BankTotalData `json:"data"`
	BaseResponse
}

type BankPageData struct {
	Metadata Metadata      `json:"metadata"`
	Items    []models.Bank `json:"items"`
}

type BankPageResponse struct {
	Data BankPageData `json:"data"`
	BaseResponse
}

type SenderIsBankResponse struct {
	Data bool `json:"data"`
	BaseResponse
}

type GetAvailablePlatformsResponse struct {
	Data string `json:"data"`
	BaseResponse
}

type CurrencyListResponse struct {
	Data []models.Currency `json:"data"`
	BaseResponse
}

type CurrencyPageData struct {
	Metadata Metadata          `json:"metadata"`
	Items    []models.Currency `json:"items"`
}

type CurrencyPageResponse struct {
	Data CurrencyPageData `json:"data"`
	BaseResponse
}

type CurrencyResponse struct {
	Data models.Currency `json:"data"`
	BaseResponse
}

type AccountBalanceResponse struct {
	Data AccountBalanceData `json:"data"`
	BaseResponse
}

type AccountBalanceData struct {
	Items []models.AmountOfBank `json:"items"`
	Total int64                 `json:"total"`
}

type BankBalanceData struct {
	Liabilities []models.AmountOfBank `json:"liabilities"`
	Claims      []models.AmountOfBank `json:"claims"`
	Issue       int64                 `json:"issue"`
	IssueLimit  int64                 `json:"issueLimit"` // TODO убрать если не быдет использоваться
}

type BankBalanceResponse struct {
	Data BankBalanceData `json:"data"`
	BaseResponse
}

type ClearingResultResponse struct {
	Data models.ClearingData `json:"data"`
	BaseResponse
}

type ClearingListResponse struct {
	Data []models.ClearingData `json:"data"`
	BaseResponse
}

type ClaimsListResponse struct {
	Data []models.ClaimsItem `json:"data"`
	BaseResponse
}

type ClearingPageData struct {
	Metadata Metadata              `json:"metadata"`
	Items    []models.ClearingData `json:"items"`
}

type ClearingPageResponse struct {
	Data ClearingPageData `json:"data"`
	BaseResponse
}

type ClaimsPageData struct {
	Metadata Metadata            `json:"metadata"`
	Items    []models.ClaimsItem `json:"items"`
}

type ClaimsPageResponse struct {
	Data ClaimsPageData `json:"data"`
	BaseResponse
}

type BankClaimsLiabilities struct {
	Claims      map[string]int64 `json:"claims"`
	Liabilities map[string]int64 `json:"liabilities"`
}

type BankClaimsLiabilitiesResponse struct {
	Data BankClaimsLiabilities `json:"data"`
	BaseResponse
}

type ClientBankItemResponse struct {
	Address         string            `json:"address"`
	BankDisplayName string            `json:"bankDisplayName"`
	State           state_enum.State  `json:"state"`
	CountryCode     string            `json:"countryCode"`
	Params          map[string]string `json:"params"`
	Owner           string            `json:"owner"`
}

type ClientBankResponse struct {
	Data ClientBankItemResponse `json:"data"`
	BaseResponse
}

type ClientBankPageData struct {
	Metadata Metadata                 `json:"metadata"`
	Items    []ClientBankItemResponse `json:"items"`
}

type ClientBankPageResponse struct {
	Data ClientBankPageData `json:"data"`
	BaseResponse
}

type ClientBankParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CustomerPageData struct {
	Metadata Metadata          `json:"metadata"`
	Items    []models.Customer `json:"items"`
}

type CustomerPageResponse struct {
	Data CustomerPageData `json:"data"`
	BaseResponse
}

type CustomersListResponse struct {
	Data []models.Customer `json:"data"`
	BaseResponse
}

type CustomerResponse struct {
	Data models.Customer `json:"data"`
	BaseResponse
}

type RoutesItem struct {
	Commission *models.CurrencyContractRoutingItem
	Middle     *models.CurrencyContractRoutingItem
	Input      *models.CurrencyContractRoutingItem
}

type ContractResponse struct {
	Data models.CurrencyExchangeContract `json:"data"`
	BaseResponse
}

type CurrencyExchangeContractPageData struct {
	Metadata Metadata                          `json:"metadata"`
	Items    []models.CurrencyExchangeContract `json:"items"`
}

type CurrencyExchangeContractPageResponse struct {
	Data CurrencyExchangeContractPageData `json:"data"`
	BaseResponse
}

type GetBestRoutesResponse struct {
	Data [][]models.CurrencyContractRoutingItem `json:"data"`
	BaseResponse
}

type WithdrawResultResponse struct {
	Data *models.WithdrawResult `json:"data"`
	BaseResponse
}

type WithdrawConfirmResponse struct {
	History models.TransactionHistory `json:"history"`
	Data    string                    `json:"data"`
	BaseResponse
}

type WithdrawRejectResponse WithdrawConfirmResponse

type TransferResponse struct {
	History models.TransactionHistory `json:"history"`
	Data    string                    `json:"data"`
	BaseResponse
}

type TransferBatchResponse struct {
	Histories []models.TransactionHistory `json:"histories"`
	Data      string                      `json:"data"`
	BaseResponse
}

type TopupResponse struct {
	Data string `json:"data"`
	BaseResponse
}

type ExecutedTransactionCurrencyExchangeContractData struct {
	Transactions         []models.ExecutedTransactionCurrencyExchangeContractItem `json:"transactions"`
	Commission           int64                                                    `json:"commission"`
	AmountInput          int64                                                    `json:"amountInput"`
	AmountOutput         int64                                                    `json:"amountOutput"`
	TransactionHistories []models.TransactionHistory                              `json:"transactionHistories"`
}

type ExecutedTransactionCurrencyExchangeContractResponse struct {
	Data ExecutedTransactionCurrencyExchangeContractData `json:"data"`
	BaseResponse
}

type CrossTransactionHistoryPageData struct {
	Metadata Metadata                         `json:"metadata"`
	Items    []models.CrossTransactionHistory `json:"items"`
}

type CrossTransactionHistoryPageResponse struct {
	Data CrossTransactionHistoryPageData `json:"data"`
	BaseResponse
}

type AccountAddressResponse struct {
	Data string `json:"data"`
	BaseResponse
}

type GetBankBalanceTotalResponseData struct {
	Bank        models.BankInfo       `json:"bank"`
	Liabilities []models.AmountOfBank `json:"liabilities"`
	Claims      []models.AmountOfBank `json:"claims"`
	Issue       int64                 `json:"issue"`
	IssueLimit  int64                 `json:"issueLimit"`
}

type GetBankBalanceTotalResponse struct {
	Data []GetBankBalanceTotalResponseData `json:"data"`
	BaseResponse
}

type InvoiceResponse struct {
	Data models.Invoice `json:"data"`
	BaseResponse
}

type InvoiceListResponse struct {
	Data []models.Invoice `json:"data"`
	BaseResponse
}

type InvoicePageData struct {
	Metadata Metadata         `json:"metadata"`
	Items    []models.Invoice `json:"items"`
}

type InvoicePageResponse struct {
	Data InvoicePageData `json:"data"`
	BaseResponse
}

type LimitData struct {
	Value int64 `json:"value"`
}

type LimitResponse struct {
	Data LimitData `json:"data"`
	BaseResponse
}

type AccountLimitsData struct {
	OperationLimit int64 `json:"operationLimit"`
	BalanceLimit   int64 `json:"balanceLimit"`
	DailyLimit     int64 `json:"dailyLimit"`
	MonthlyLimit   int64 `json:"monthlyLimit"`
}

type AccountLimitsResponse struct {
	Data AccountLimitsData `json:"data"`
	BaseResponse
}

type SafeDealResponse struct {
	Data models.DealResponseData `json:"data"`
	BaseResponse
}

type SafeDealDepositResponse struct {
	Data models.SafeDealDeposit `json:"data"`
	BaseResponse
}
