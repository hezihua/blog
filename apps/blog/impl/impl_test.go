package impl_test

import (
	"context"
	"hezihua/apps"
	"hezihua/apps/blog"
	"hezihua/test/tools"
	"testing"
)

var (
	impl blog.Service
	ctx = context.Background()
)

func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Author = "author"
	req.Title = "title"
	req.Content = "content"
	req.Summary = "summary"
	ins, err := impl.CreateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ins)
}

func TestQueryBlog(t *testing.T) {
	req := blog.NewQueryBlogRequest()
	// req.Keywords = "title"
	// req.Status = blog.StatusPublished
	set, err := impl.QueryBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(set)
}




func init() {
	tools.DevelopmentSet()
	impl = apps.GetInternalApp(blog.AppName).(blog.Service)
}