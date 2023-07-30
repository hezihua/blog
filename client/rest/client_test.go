package rest_test

import (
	"context"
	"fmt"
	"hezihua/apps/blog"
	"hezihua/client/rest"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
)

var (
	client *rest.Client
	ctx 	= context.Background()
)

func TestCreateBlog(t *testing.T) {
	// err := client.Login() 
	// if err != nil {
	// 	panic(err)
	// }
	req := blog.NewCreateBlogRequest()
	req.Author = "author1"
	req.Title = "title2"
	req.Content = "content3"
	req.Summary = "summary4"
	client.CreateBlog(ctx, req)
}

// func TestLogin(t *testing.T) {
// 	err := client.Login() 
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//   t.Log(client.Session())
// }




func init(){
	zap.DevelopmentSetup()
	conf := rest.NewDefaultConfig()	
	client = rest.NewClient(conf)
	err := client.Login() 
	if err != nil {
		panic(err)
	}
  fmt.Println(client.Session())
}