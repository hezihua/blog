package middleware

import (
	"fmt"
	"hezihua/apps"
	"hezihua/apps/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthFunc(c *gin.Context) {
  ck, err := c.Request.Cookie("username")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": err.Error(),
		})
		c.Abort()
		return
	}
	username := ck.Value
	ck, err = c.Request.Cookie("session")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": err.Error(),
		})
		c.Abort()
		return
	}
	session := ck.Value
	svc := apps.GetInternalApp(user.AppName).(user.Service)
	err = svc.CheckIsLogin(c.Request.Context(), user.NewCheckIsLoginRequest(username, session))
	fmt.Println(ck)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": err.Error(),
		})
		c.Abort()
		return
	}

	c.Next()




}