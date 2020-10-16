package account_type_enum

type AccountType int

const (
	Undefined AccountType = iota
	Client
	Bank
	Agent
)

func (accountType AccountType) Str() string {
	return [...]string{"Undefined", "Client", "Bank", "Agent"}[accountType]
}

func GetAccountTypes() []AccountType {
	return []AccountType{Undefined, Client, Bank, Agent}
}
