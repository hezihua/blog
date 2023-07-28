package impl

import (
	"fmt"
	"hezihua/apps"
	"hezihua/apps/blog"
)

type impl struct {
	blog blog.Service
}


func (i *impl) Init() error {
	fmt.Println(blog.AppName)
	fmt.Println(apps.GetInternalApp(blog.AppName))
	i.blog = apps.GetInternalApp(blog.AppName).(blog.Service)
	return nil
}

func (i *impl) Name() string {
	return blog.AppName
}




func init(){
	apps.Registry(&impl{})
}