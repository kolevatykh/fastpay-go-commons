package account_type

type AccountType int

const (
	Undefined AccountType = iota
	Client
	Bank
)

func (accountType AccountType) String() string {
	return [...]string{"Undefined", "Client", "Bank"}[accountType]
}
