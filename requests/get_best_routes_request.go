package requests

type GetBestRoutesRequest struct {
	Amount             int64  `json:"amount" validate:"required"`
	CurrencyCodeFrom   int    `json:"currencyCodeFrom" validate:"required,min=0"`
	CurrencyCodeTo     int    `json:"currencyCodeTo" validate:"required,min=0"`
	CustomerIdentifier string `json:"customerIdentifier" validate:"omitempty,validHex64"`
	CountryCode        string `json:"countryCode" validate:"omitempty,min=3,max=3"`
	To                 string `json:"to" validate:"omitempty,validHex40or64"`
}
