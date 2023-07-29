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
	ins, err := impl.CreateBlog(ctx, &blog.CreateBlogRequest{})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ins)
}

func init() {
	tools.DevelopmentSet()
	impl = apps.GetInternalApp(blog.AppName).(blog.Service)
}