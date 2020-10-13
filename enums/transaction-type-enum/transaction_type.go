package transaction_type_enum

type TransactionType int

const (
	Undefined TransactionType = iota
	Topup
	Transfer
	Withdraw
	Payment
	InstantPayment
)

func (transactionType TransactionType) String() string {
	return [...]string{"Undefined", "Topup", "Transfer", "Withdraw", "Payment", "InstantPayment"}[transactionType]
}
