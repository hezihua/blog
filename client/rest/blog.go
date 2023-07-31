package rest

import (
	"context"
	"hezihua/apps/blog"
)




func (c *Client) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (
	*blog.Blog, error) {
		// resp, _ := http.Post("http://localhost:8080/api/v1/blog", "application/json", req)
		// body, _ := io.ReadAll(resp.Body)
		ins := &blog.Blog{}
		err := c.c.Post("blogs").Body(in).Do(ctx).Into(ins)
		if err != nil {
			return nil, err
		}
		return nil, nil

}
