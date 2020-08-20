package requests

type TransferBatchRequest struct {
	Transfers []TransferRequest `json:"transfers" valid:"required"`
}
