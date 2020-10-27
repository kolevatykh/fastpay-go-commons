package requests

type SetAvailablePlatformsRequest struct {
	TechnicalSignRequest
	AvailablePlatforms string `json:"availablePlatforms" valid:"required"`
}
