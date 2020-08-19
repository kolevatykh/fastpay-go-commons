package requests

type GetCustomerRequest struct {
	BankId      string `json:"bankId" valid:"required~60364"`
	Identifier  string `json:"identifier" valid:"required~60304,validHex64~60303"`
	CountryCode string `json:"countryCode" valid:"required~60365,stringlength(3|3)"`
}
