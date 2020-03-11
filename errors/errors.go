package errors

type BaseError struct {
	Code    uint8  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
