package impl

import (
	"hezihua/apps"
	"hezihua/apps/blog"
	"hezihua/apps/tag"
	"hezihua/conf"

	"gorm.io/gorm"
)



type impl struct {
	blog blog.Service
	db	 *gorm.DB
}

// 注入依赖
func (i *impl) Init() error {
	i.blog = apps.GetInternalApp(blog.AppName).(blog.Service)
	i.db = conf.C().Mysql.ORM().Debug()
	return nil
}

func (i *impl) Name() string {
	return tag.AppName
}


func init(){
	apps.Registry(&impl{})
}