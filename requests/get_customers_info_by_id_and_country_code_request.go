package requests

type GetCustomersInfoByIdAndCountryCodeRequest struct {
	Identifier  string `json:"identifier" validate:"required,validHex64"`
	CountryCode string `json:"countryCode" validate:"required,min=3,max=3"`
}
