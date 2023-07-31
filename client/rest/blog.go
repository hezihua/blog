package rest

import (
	"context"
	"fmt"
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

func (c *Client) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (
	*blog.BlogSet, error) {
		// resp, _ := http.Post("http://localhost:8080/api/v1/blog", "application/json", req)
		// body, _ := io.ReadAll(resp.Body)
		set := &blog.BlogSet{}
		err := c.c.Get("blogs").
		Param("page_size", fmt.Sprintf("%d", in.PageSize)).
		Param("page_number", fmt.Sprintf("%d", in.PageNumber)).
		Do(ctx).
		Into(set)
		if err != nil {
			return nil, err
		}
		return set, nil

}

