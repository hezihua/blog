package api

import (
	"fmt"
	"hezihua/apps/tag"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateTag(c *gin.Context) {
	req := tag.NewDefaultCreateTagRequest()
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(req, "11111")
	ins, err := h.svc.CreateTag(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ins)
}