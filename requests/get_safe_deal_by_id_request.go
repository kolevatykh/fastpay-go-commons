package requests

type GetSafeDealByIdRequest struct {
	Id string `json:"id" valid:"required"`
}
