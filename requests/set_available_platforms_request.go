package requests

type SetAvailablePlatformsRequest struct {
	AvailablePlatforms string `json:"availablePlatforms" validate:"required"`
}
