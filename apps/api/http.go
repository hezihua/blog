package api

import (
	"hezihua/apps/user"

	"github.com/gin-gonic/gin"
)

func NewHandler(svc user.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}
type Handler struct {

	svc user.Service
}

// Registry 注册路由
func (h *Handler) Registry(r gin.IRouter) {

	r.POST("/login", h.Login)
	r.POST("/logout", h.Logout)

}