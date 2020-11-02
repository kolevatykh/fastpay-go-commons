package requests

type GetCustomerRequest struct {
	BankAddress string `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
	Identifier  string `json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64~ErrorIdentifierNotFolowingRegex"`
	CountryCode string `json:"countryCode" valid:"required~ErrorCountryCodeNotPassed,stringlength(3|3)"`
}
