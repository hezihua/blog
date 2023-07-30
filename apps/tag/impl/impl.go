package impl

import (
	"hezihua/apps"
	"hezihua/apps/blog"
	"hezihua/apps/tag"
)



type impl struct {
	blog blog.Service
}

// 注入依赖
func (i *impl) Init() error {
	i.blog = apps.GetInternalApp(blog.AppName).(blog.Service)
	return nil
}

func (i *impl) Name() string {
	return tag.AppName
}


func init(){
	apps.Registry(&impl{})
}