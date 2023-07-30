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

func TestDescribeBlog(t *testing.T) {
	req := blog.NewDescribeBlogRequest(3)
	ins, err := impl.DescribeBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ins)
}

func TestQueryBlog(t *testing.T) {
	req := blog.NewQueryBlogRequest()
	// req.Keywords = "test"
	// req.Status = blog.StatusPublished
	set, err := impl.QueryBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(set)
}


func TestPutUpdateBlog(t *testing.T) {
	data := blog.NewCreateBlogRequest()
	req := blog.NewPutUpdate(3,data)
	// req.Keywords = "test"
	// req.Status = blog.StatusPublished
	t.Log(req)
	ins, err := impl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ins)
}

func TestPatchUpdateBlog(t *testing.T) {
	data := blog.NewCreateBlogRequest()
	req := blog.NewPatchUpdate(3,data)
	req.Title = "123"
	// req.Keywords = "test"
	// req.Status = blog.StatusPublished
	t.Log(req)
	ins, err := impl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ins)
}

func TestDeleteBlog(t *testing.T) {
	req := blog.NewDeleteBlogRequest(3)
	t.Log(req)
	ins, err := impl.DeleteBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ins)
}




func init() {
	tools.DevelopmentSet()
	impl = apps.GetInternalApp(blog.AppName).(blog.Service)
}