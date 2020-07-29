package currency_type_enum

type CurrencyType int

const (
	AccountDigitalUnit CurrencyType = iota
	DigitalAsset
)

func (currencyType CurrencyType) String() string {
	return [...]string{"AccountDigitalUnit", "DigitalAsset"}[currencyType]
}

func Parse(stringJuridicalType string) CurrencyType {
	switch stringJuridicalType {
	case "AccountDigitalUnit":
		return AccountDigitalUnit
	case "DigitalAsset":
		return DigitalAsset
	default:
		return AccountDigitalUnit
	}
}

func GetCurrencyTypes() []CurrencyType {
	return []CurrencyType{AccountDigitalUnit, DigitalAsset}
}
