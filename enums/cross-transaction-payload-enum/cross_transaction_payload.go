package cross_transaction_payload_enum

type CrossTransactionPayload string

const (
	Undefined  CrossTransactionPayload = ""
	Sell       CrossTransactionPayload = "sell"
	Buy        CrossTransactionPayload = "buy"
	Commission CrossTransactionPayload = "commission"
	Withdraw   CrossTransactionPayload = "withdraw"
)
