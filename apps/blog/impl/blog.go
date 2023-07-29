package impl

import (
	"context"
	"hezihua/apps/blog"
)

func (i *impl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	return nil, nil
}
func (i *impl) QueryBlog(ctx context.Context, req *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	return nil, nil
}
func (i *impl) DescribeBlog(context.Context, *blog.DescribeBlogRequest) (*blog.Blog, error) {
	return nil, nil
}
func (i *impl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	return nil, nil
}
func (i *impl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

