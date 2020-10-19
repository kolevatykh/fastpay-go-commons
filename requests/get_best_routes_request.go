package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contracts-type-enum"

type GetBestRoutesRequest struct {
	Amount             int64                                                               `json:"amount" valid:"required~ErrorAmountNotPassed"`
	CurrencyCodeFrom   int                                                                 `json:"currencyCodeFrom" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	CurrencyCodeTo     int                                                                 `json:"currencyCodeTo" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	CustomerIdentifier string                                                              `json:"customerIdentifier" valid:"validHex64~ErrorIdentifierNotFolowingRegex"`
	CountryCode        string                                                              `json:"countryCode" valid:"stringlength(3|3)"`
	To                 string                                                              `json:"to" valid:"validHex40or64~ErrorAddressOrIdentifierNotFolowingRegex"`
	Type               currency_exchange_contracts_type_enum.CurrencyExchangeContractsType `json:"type"`
}

func (getBestRoutes *GetBestRoutesRequest) SetDefaults() {
	if getBestRoutes.Type == currency_exchange_contracts_type_enum.Undefined {
		getBestRoutes.Type = currency_exchange_contracts_type_enum.CurrencyExchangeContract
	}
}
