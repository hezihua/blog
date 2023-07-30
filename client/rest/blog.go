package rest

import (
	"context"
	"hezihua/apps/blog"

	"github.com/gin-gonic/gin"
)

func (c *Client) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (
	*blog.Blog, error) {
		// resp, _ := http.Post("http://localhost:8080/api/v1/blog", "application/json", req)
		// body, _ := io.ReadAll(resp.Body)
		ins := &blog.Blog{}
		err := c.c.Post("blogs").Body(in).Do(ctx).Into(gin.H{
			"data": ins,
		})
		if err != nil {
			return nil, err
		}
		return nil, nil

}
