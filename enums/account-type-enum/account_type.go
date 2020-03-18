package account_type_enum

type AccountType int

const (
	Undefined AccountType = iota
	Client
	Bank
)

func (accountType AccountType) String() string {
	return [...]string{"Undefined", "Client", "Bank"}[accountType]
}
