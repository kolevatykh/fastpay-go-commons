package responses

type BaseResponse struct {
	ErrorCode     uint   `json:"errorCode"`
	TransactionId uint   `json:"transactionId"`
	SenderAddress string `json:"senderAddress"`
}
