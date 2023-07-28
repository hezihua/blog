package blog

type Blog struct {
	Id       int64 `json:"id"`
	CreateAt int64 `json:"create_at"`
	*CreateBlogRequest
}