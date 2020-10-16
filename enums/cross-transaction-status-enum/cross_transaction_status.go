package cross_transaction_status_enum

type CrossTransactionStatus int

const (
	Undefined CrossTransactionStatus = iota
	Sent
	WaitWithdraw
	WithdrawToHlfSuspended
	ProgressWithdraw
	WithdrawInAbsSuspended
	RejectedWithdraw
	Executed
	Error
)

func (crossTransactionStatus CrossTransactionStatus) Str() string {
	return [...]string{"Undefined", "Sent", "WaitWithdraw", "WithdrawToHlfSuspended", "ProgressWithdraw",
		"WithdrawInAbsSuspended", "RejectedWithdraw", "Executed", "Error"}[crossTransactionStatus]
}
