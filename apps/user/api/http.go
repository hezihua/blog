package api

import (
	"hezihua/apps"
	"hezihua/apps/user"

	"github.com/gin-gonic/gin"
)

func Newhandler(svc user.Service) *handler {
	return &handler{
		svc: svc,
	}
}
type handler struct {

	svc user.Service
}

// Registry 注册路由
func (h *handler) RegistryHandler(r gin.IRouter) {

	r.POST("/login", h.Login)
	r.POST("/logout", h.Logout)

}

func (h *handler) Init() error {
	h.svc = apps.GetInternalApp(user.AppName).(user.Service)
	return nil
}

func (h *handler) Name() string {
	return user.AppName
}

func init() {
	apps.RegistryHttp(&handler{})
}