package requests

type CreateCustomerRequest struct {
	TechnicalSignRequest
	Identifier          string `json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64~ErrorIdentifierNotFolowingRegex"`
	CustomerDisplayName string `json:"customerDisplayName"`
	BankAddress         string `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}
