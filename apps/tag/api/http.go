package api

import (
	"hezihua/apps"
	"hezihua/apps/tag"
	"hezihua/protocol/middleware"

	"github.com/gin-gonic/gin"
)

func Newhandler(svc tag.Service) *handler {
	return &handler{
		svc: svc,
	}
}
type handler struct {

	svc tag.Service
}

// Registry 注册路由
func (h *handler) RegistryHandler(r gin.IRouter) {
	r.Use(middleware.AuthFunc)
	r.POST("/", h.CreateTag)

}

func (h *handler) Init() error {
	h.svc = apps.GetInternalApp(tag.AppName).(tag.Service)
	return nil
}

func (h *handler) Name() string {
	return tag.AppName
}

func init() {
	apps.RegistryHttp(&handler{})
}