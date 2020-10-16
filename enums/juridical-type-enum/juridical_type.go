package juridical_type_enum

type JuridicalType int

const (
	Undefined JuridicalType = iota
	Individual
	IndividualEntrepreneur
	LegalEntity
)

func (juridicalType JuridicalType) Str() string {
	return [...]string{"Undefined", "Individual", "IndividualEntrepreneur", "LegalEntity"}[juridicalType]
}

func Parse(stringJuridicalType string) JuridicalType {
	switch stringJuridicalType {
	case "Individual":
		return Individual
	case "IndividualEntrepreneur":
		return IndividualEntrepreneur
	case "LegalEntity":
		return LegalEntity
	default:
		return Undefined
	}
}

func GetAllJuridicalTypes() []JuridicalType {
	return []JuridicalType{Undefined, Individual, IndividualEntrepreneur, LegalEntity}
}
