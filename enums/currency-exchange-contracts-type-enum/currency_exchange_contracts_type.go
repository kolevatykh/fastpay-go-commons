package currency_exchange_contracts_type_enum

type CurrencyExchangeContractsType int

const (
	Undefined CurrencyExchangeContractsType = iota
	CurrencyExchangeContract
	OfferExchangeContract
)

func (currencyExchangeContractsType CurrencyExchangeContractsType) Str() string {
	return [...]string{"Undefined", "CurrencyExchangeContract", "OfferExchangeContract"}[currencyExchangeContractsType]
}

func Parse(stringCurrencyExchangeContractsType string) CurrencyExchangeContractsType {
	switch stringCurrencyExchangeContractsType {
	case "CurrencyExchangeContract":
		return CurrencyExchangeContract
	case "OfferExchangeContract":
		return OfferExchangeContract
	default:
		return Undefined
	}
}

func GetCurrencyExchangeContractsTypes() []CurrencyExchangeContractsType {
	return []CurrencyExchangeContractsType{Undefined, CurrencyExchangeContract, OfferExchangeContract}
}
