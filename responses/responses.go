package responses

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

type BankResponse struct {
	Data models.Bank `json:"data"`
	BaseResponse
}

type AccountResponse struct {
	Data models.Account `json:"data"`
	BaseResponse
}

type AccountListResponse struct {
	Data []models.Account `json:"data"`
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
	Items []models.AmountOfBank `json:"items"`
	Total int64                 `json:"total"`
}

type WithdrawResultResponse struct {
	Data models.WithdrawResult `json:"data"`
	BaseResponse
}

type WithdrawConfirmResponse struct {
	Data models.TransactionHistory `json:"data"`
	BaseResponse
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

type WithdrawRejectResponse WithdrawConfirmResponse
