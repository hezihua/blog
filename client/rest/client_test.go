package rest_test

import (
	"context"
	"hezihua/apps/blog"
	"hezihua/client/rest"
	"testing"
)

var (
	client *rest.Client
	ctx 	= context.Background()
)

func TestXxx(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Author = "author1"
	req.Title = "title2"
	req.Content = "content3"
	req.Summary = "summary4"
	client.CreateBlog(ctx, req)
}




func init(){
	conf := rest.NewDefaultConfig()	
	client = rest.NewClient(conf)
}