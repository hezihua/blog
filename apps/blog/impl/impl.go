package impl

import (
	"hezihua/apps"
	"hezihua/apps/blog"
	"hezihua/apps/tag"
	"hezihua/conf"

	"gorm.io/gorm"
)

type impl struct {
	tag tag.Service
	// 连接池
	db *gorm.DB
}


func (i *impl) Init() error {
	i.tag = apps.GetInternalApp(tag.AppName).(tag.Service)

	i.db = conf.C().Mysql.ORM().Debug()
	return nil
}

func (i *impl) Name() string {
	return blog.AppName
}




func init(){
	apps.Registry(&impl{})
}