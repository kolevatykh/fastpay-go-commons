package requests

type CreateCustomerRequest struct {
	Identifier          string `json:"identifier" valid:"required~60304,validHex64~60367"`
	CustomerDisplayName string `json:"customerDisplayName"`
}
