package impl

import (
	"hezihua/apps"
	"hezihua/apps/tag"
)



type impl struct {
	tag tag.Service
}

// 注入依赖
func (i *impl) Init() error {
	i.tag = apps.GetInternalApp(tag.AppName).(tag.Service)
	return nil
}

func (i *impl) Name() string {
	return tag.AppName
}


func init(){
	apps.Registry(&impl{})
}