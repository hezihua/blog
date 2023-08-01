package blog_test

import (
	"encoding/json"
	"hezihua/apps/blog"
	"testing"
)

func TestStatusMarshaler(t *testing.T) {
	// dj, err :=blog.StatusDraft.MarshalJSON()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(dj)
	b := &blog.Blog{CreateBlogRequest: blog.NewCreateBlogRequest()}
	dj, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(dj))
}
func TestStatusUnmarshaler(t *testing.T){
	// var s blog.Status
	// err := s.UnmarshalJSON([]byte(`"draft"`))
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(s)
	dataJ := `
	{
		"title": "test",
		"author": "test",
		"content": "test",
		"summary": "test",
		"status": "published"
	}
	`
	b := blog.NewBlog(blog.NewCreateBlogRequest())

	err := json.Unmarshal([]byte(dataJ),b)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(b)
}