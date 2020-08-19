package requests

type GetBestRoutesRequest struct {
	Amount             int64  `json:"amount" valid:"required~60313"`
	CurrencyCodeFrom   int    `json:"currencyCodeFrom" valid:"required~60318,range(0|999)~60317"`
	CurrencyCodeTo     int    `json:"currencyCodeTo" valid:"required~60318,range(0|999)~60317"`
	CustomerIdentifier string `json:"customerIdentifier" valid:"validHex64~60303"`
	CountryCode        string `json:"countryCode" valid:"stringlength(3|3)"`
	To                 string `json:"to" valid:"validHex40or64~60367"`
}
