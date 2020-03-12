package juridical_type

type JuridicalType int

const (
	Undefined JuridicalType = iota
	Individual
	IndividualEntrepreneur
	LegalEntity
)

func (juridicalType JuridicalType) String() string {
	return [...]string{"Undefined", "Individual", "IndividualEntrepreneur", "LegalEntity"}[juridicalType]
}
