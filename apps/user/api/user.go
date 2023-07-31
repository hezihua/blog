package api

import (
	"hezihua/apps/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Login(c *gin.Context) {
	req := user.NewLoginRequestFromBasicAuth(c.Request)
	sess, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"code": 	http.StatusUnauthorized,
		})
		return
	}

	// 设置cookie
	c.SetCookie("session", sess.Id, 3600, "", "", false, true)
	c.SetCookie("username", sess.Username, 3600, "", "", false, true)


	c.JSON(http.StatusOK, gin.H{
		"message": sess,
		"code":    http.StatusOK,
	})
}

func (h *handler) Logout(c *gin.Context) {	
	req := user.NewLogoutRequest()
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code": 	http.StatusBadRequest,
		})
	}
  err = h.svc.Logout(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code": 	http.StatusInternalServerError,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "logout success",
		"code":    http.StatusOK,
	})
}
