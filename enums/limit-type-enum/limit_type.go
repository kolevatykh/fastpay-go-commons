package limit_type_enum

type LimitType int

const (
	Undefined LimitType = iota
	Operation
	Daily
	Monthly
	Balance
)

func (limitType LimitType) Str() string {
	return [...]string{"Undefined", "Operation", "Daily", "Monthly", "Balance"}[limitType]
}

func Parse(stringLimitType string) LimitType {
	switch stringLimitType {
	case "Operation":
		return Operation
	case "Daily":
		return Daily
	case "Monthly":
		return Monthly
	case "Balance":
		return Balance
	default:
		return Undefined
	}
}

func GetAllLimitTypes() []LimitType {
	return []LimitType{Undefined, Operation, Daily, Monthly, Balance}
}
