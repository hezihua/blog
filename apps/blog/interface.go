package blog

import "context"

var (
	AppName = "blog"
)

type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)

	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)

	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)

  UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)

	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)

}


type CreateBlogRequest struct {
  Title string `json:"title"`
	Author string `json:"author"`
	Content string `json:"content"`
	Summary string `json:"summary"`
	Status Status `json:"status"`
}



type QueryBlogRequest struct {
	PageSize uint `json:"page_size"`
	PageNumber uint `json:"page_num"`
	Keywords string `json:"keywords"`
	// 这里传指针
	Status *Status `json:"status"`

}

type BlogSet struct {
	Total int64 `json:"total"`
	Items []*Blog `json:"items"`
}

func (s *BlogSet) AddItem(item *Blog) {
	s.Items = append(s.Items, item)
}

type DescribeBlogRequest struct {
	Id int `json:"id"`

}

type UpdateBlogRequest struct {
	Id int `json:"id"`
	UpdateMode UpdateMode `json:"update_mode"`
	*CreateBlogRequest
}

type DeleteBlogRequest struct {

}