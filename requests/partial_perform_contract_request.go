package requests

type PartialPerformContractRequest struct {
	PerformContractRequest
	ActualAmountInitiator int64 `json:"actualAmountInitiator" valid:"required,range(0|9223372036854775807)~ErrorAmountNegative"`
	ActualAmountAcceptor  int64 `json:"actualAmountAcceptor" valid:"required,range(0|9223372036854775807)~ErrorAmountNegative"`
}
