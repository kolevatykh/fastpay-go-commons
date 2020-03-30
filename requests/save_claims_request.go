package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type SaveClaimsRequest struct {
	Claims []models.ClaimsItem `json:"claims" validate:"required,dive"`
}
