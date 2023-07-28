package blog

import "context"

var (
	AppName = "blog"
)

type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogResponse, error)
}


type CreateBlogRequest struct {

}

type CreateBlogResponse struct {

}