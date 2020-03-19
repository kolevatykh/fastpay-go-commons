package requests

type GetPageRequest struct {
	PageSize int32  `json:"pageSize" validate:"required"`
	Bookmark string `json:"bookmark" validate:"required"`
}
