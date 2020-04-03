package requests

type CreateClientBankRequest struct {
	BankDisplayName string            `json:"bankDisplayName" validate:"required"`
	CountryCode     string            `json:"countryCode"`
	Params          map[string]string `json:"params"`
}
