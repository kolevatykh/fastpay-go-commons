package requests

type CreateCustomerRequest struct {
	Identifier          string `json:"identifier" validate:"required,validHex64"`
	CustomerDisplayName string `json:"customerDisplayName"`
}
