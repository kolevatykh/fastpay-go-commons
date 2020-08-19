package requests

type GetBestRoutesRequest struct {
	Amount             int64  `json:"amount" valid:"required"`
	CurrencyCodeFrom   int    `json:"currencyCodeFrom" valid:"required,min(0)"`
	CurrencyCodeTo     int    `json:"currencyCodeTo" valid:"required,min(0)"`
	CustomerIdentifier string `json:"customerIdentifier" valid:"validHex64"`
	CountryCode        string `json:"countryCode" valid:"min(3),max(3)"`
	To                 string `json:"to" valid:"validHex40or64"`
}
