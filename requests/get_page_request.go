package requests

type GetPageRequest struct {
	PageSize int32  `json:"pageSize" valid:"required"`
	Bookmark string `json:"bookmark"`
}
