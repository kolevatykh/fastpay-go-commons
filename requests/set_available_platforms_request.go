package requests

type SetAvailablePlatformsRequest struct {
	AvailablePlatforms string `json:"availablePlatforms" valid:"required"`
}
