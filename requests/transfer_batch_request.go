package requests

type TransferBatchRequest struct {
	Transfers []TransferRequest `json:"transfers" valid:"required"`
}

func (transferRequest *TransferBatchRequest) SetDefaults() {
	for i, _ := range transferRequest.Transfers {
		transferRequest.Transfers[i].SetDefaults()
	}
}
