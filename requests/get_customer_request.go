package requests

type GetCustomerRequest struct {
	BankId      string `json:"bankId" validate:"required"`
	Identifier  string `json:"identifier" validate:"required,validHex64"`
	CountryCode string `json:"countryCode" validate:"required,min=3,max=3"`
}
