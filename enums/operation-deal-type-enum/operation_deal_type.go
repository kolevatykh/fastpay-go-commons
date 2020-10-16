package operation_deal_type_enum

type OperationDealType int

const (
	Undefined OperationDealType = iota
	Transfer
	Deposit
)

func (operationDealType OperationDealType) Str() string {
	return [...]string{"Undefined", "Transfer", "Deposit"}[operationDealType]
}

func Parse(operationDealType string) OperationDealType {
	switch operationDealType {
	case "Transfer":
		return Transfer
	case "Deposit":
		return Deposit
	default:
		return Undefined
	}
}

func GetAllOperationDealTypes() []OperationDealType {
	return []OperationDealType{Undefined, Transfer, Deposit}
}
