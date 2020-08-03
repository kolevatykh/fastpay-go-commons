package responses

type BaseResponse struct {
	ErrorCode     int    `json:"errorCode"`
	TransactionId string `json:"transactionId"`
	SenderAddress string `json:"senderAddress"`
}
