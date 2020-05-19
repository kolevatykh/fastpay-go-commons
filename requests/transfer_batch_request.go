package requests

type TransferBatchRequest struct {
	Transfers []TransferRequest `json:"transfers" validate:"required,dive"`
}
