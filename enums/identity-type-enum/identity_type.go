package identity_type_enum

type IdentityType int

const (
	Undefined IdentityType = iota
	None
	Simple
	Identified
)

func (identityType IdentityType) Str() string {
	return [...]string{"Undefined", "None", "Simple", "Identified"}[identityType]
}

func Parse(stringIdentityType string) IdentityType {
	switch stringIdentityType {
	case "None":
		return None
	case "Simple":
		return Simple
	case "Identified":
		return Identified
	default:
		return Undefined
	}
}

func GetAllIdentityTypes() []IdentityType {
	return []IdentityType{Undefined, None, Simple, Identified}
}
