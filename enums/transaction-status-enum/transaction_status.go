package transaction_status_enum

type TransactionStatus int

const (
	Undefined TransactionStatus = iota
	InProgress
	Rejected
	Executed
)

func (transactionStatus TransactionStatus) String() string {
	return [...]string{"Undefined", "InProgress", "Rejected", "Executed"}[transactionStatus]
}
