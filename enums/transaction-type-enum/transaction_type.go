package transaction_type_enum

type TransactionType int

const (
	Undefined TransactionType = iota
	Topup
	Transfer
	Withdraw
	Payment
	InstantPayment
	TransferToDeposit
	TransferFromDeposit
	RefundToDeposit
)

func (transactionType TransactionType) Str() string {
	return [...]string{"Undefined", "Topup", "Transfer", "Withdraw", "Payment", "InstantPayment",
		"TransferToDeposit", "TransferFromDeposit", "RefundToDeposit"}[transactionType]
}
