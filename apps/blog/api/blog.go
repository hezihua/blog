package api

import (
	"hezihua/apps/blog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateBLog(c *gin.Context) {
	req := blog.NewCreateBlogRequest()
	// io.ReadAll(c.Request.Body)
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
	}
	ins, err := h.svc.CreateBlog(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": ins,
	})
}
func (h *handler) QueryBlog(c *gin.Context) {
	// h.svc.QueryBlog()
}
func (h *handler) DescribeBLog(c *gin.Context) {
	// h.svc.DescribeBlog()
}
func (h *handler) PutBLog(c *gin.Context) {
	
}
func (h *handler) PatchBLog(c *gin.Context) {
	
}
func (h *handler) DeleteBLog(c *gin.Context) {
	// h.svc.DeleteBlog()
}