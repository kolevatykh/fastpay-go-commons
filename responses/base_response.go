package responses

type BaseResponse struct {
	ErrorCode     int    `json:"errorCode"`
	TransactionId int    `json:"transactionId"`
	SenderAddress string `json:"senderAddress"`
}
