package api

import (
	"hezihua/apps"
	"hezihua/apps/blog"

	"github.com/gin-gonic/gin"
)

func Newhandler(svc blog.Service) *handler {
	return &handler{
		svc: svc,
	}
}
type handler struct {

	svc blog.Service
}

// Registry 注册路由
func (h *handler) RegistryHandler(r gin.IRouter) {

	r.POST("/", h.CreateBLog)
	r.GET("/", h.QueryBlog)
	r.PUT("/:id", h.PutBLog)
	r.PATCH("/:id", h.PatchBLog)
	r.DELETE("/:id", h.DeleteBLog)

}

func (h *handler) Init() error {
	h.svc = apps.GetInternalApp(blog.AppName).(blog.Service)
	return nil
}

func (h *handler) Name() string {
	return blog.AppName
}

func init() {
	apps.RegistryHttp(&handler{})
}