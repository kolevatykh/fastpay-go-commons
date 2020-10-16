package currency_type_enum

type CurrencyType int

const (
	Undefined CurrencyType = iota
	DigitalCurrency
	DigitalAsset
)

func (currencyType CurrencyType) Str() string {
	return [...]string{"Undefined", "DigitalCurrency", "DigitalAsset"}[currencyType]
}

func Parse(stringCurrencyType string) CurrencyType {
	switch stringCurrencyType {
	case "DigitalCurrency":
		return DigitalCurrency
	case "DigitalAsset":
		return DigitalAsset
	default:
		return Undefined
	}
}

func GetCurrencyTypes() []CurrencyType {
	return []CurrencyType{Undefined, DigitalCurrency, DigitalAsset}
}
