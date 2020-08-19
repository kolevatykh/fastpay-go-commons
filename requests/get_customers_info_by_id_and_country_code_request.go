package requests

type GetCustomersInfoByIdAndCountryCodeRequest struct {
	Identifier  string `json:"identifier" valid:"required~60304,validHex64~60303"`
	CountryCode string `json:"countryCode" valid:"required~60365,stringlength(3|3)"`
}
