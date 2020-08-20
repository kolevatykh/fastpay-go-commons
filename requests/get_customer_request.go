package requests

type GetCustomerRequest struct {
	BankId      string `json:"bankId" valid:"required~ErrorBankIdNotPassed"`
	Identifier  string `json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64~ErrorIdentifierNotFolowingRegex"`
	CountryCode string `json:"countryCode" valid:"required~ErrorCountryCodeNotPassed,stringlength(3|3)"`
}
