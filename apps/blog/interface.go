package blog

import "context"

var (
	AppName = "blogs"
)

type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)

	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)

	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)

  UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)

	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)

}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Status: StatusDraft,
	}
}

type CreateBlogRequest struct {
  Title string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Content string `json:"content" validate:"required"`
	Summary string `json:"summary" validate:"required"`
	Status Status `json:"status"`
}

func (req *QueryBlogRequest) Offset() int64 {
	return int64(req.PageNumber - 1) * int64(req.PageSize)
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize: 10,
		PageNumber: 1,
	}
}

type QueryBlogRequest struct {
	PageSize uint `json:"page_size"`
	PageNumber uint `json:"page_num"`
	Keywords string `json:"keywords"`
	// 这里传指针
	Status *Status `json:"status"`
  Author string `json:"author"`
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type BlogSet struct {
	Total int64 `json:"total"`
	Items []*Blog `json:"items"`
}

func (s *BlogSet) AddItem(item *Blog) {
	s.Items = append(s.Items, item)
}

func NewDescribeBlogRequest(id int) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		Id: id,
	}
}

type DescribeBlogRequest struct {
	Id int `json:"id"`

}

func NewPutUpdate(id int, req *CreateBlogRequest) *UpdateBlogRequest {
	return NewUpdateBlogRequest(PUT, id, req)
}
func NewPatchUpdate(id int, req *CreateBlogRequest) *UpdateBlogRequest {
	return NewUpdateBlogRequest(PATCH, id, req)
}

func NewUpdateBlogRequest (mode UpdateMode,id int, req *CreateBlogRequest) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		UpdateMode: mode,
		Id: id,
		CreateBlogRequest: req,
	}
}

type UpdateBlogRequest struct {
	Id int `json:"id"`
	UpdateMode UpdateMode `json:"update_mode"`
	*CreateBlogRequest
}

func NewDeleteBlogRequest(id int) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		Id: id,
	}
}

type DeleteBlogRequest struct {
	Id int `json:"id"`
}