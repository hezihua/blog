package blog_test

import (
	"hezihua/apps/blog"
	"testing"
)

func TestCreateBlog(t *testing.T) {
	req := blog.CreateBlogRequest{
		Status: blog.StatusDraft,
	}
}
