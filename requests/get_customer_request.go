package requests

type GetCustomerRequest struct {
	BankId      string `json:"bankId" valid:"required"`
	Identifier  string `json:"identifier" valid:"required,validHex64"`
	CountryCode string `json:"countryCode" valid:"required,min(3),max(3)"`
}
