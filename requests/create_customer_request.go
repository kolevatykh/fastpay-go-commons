package requests

type CreateCustomerRequest struct {
	Identifier          string `json:"identifier" valid:"required,validHex64"`
	CustomerDisplayName string `json:"customerDisplayName"`
}
