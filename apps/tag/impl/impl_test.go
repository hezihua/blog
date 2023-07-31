package impl_test

import (
	"context"
	"hezihua/apps"
	"hezihua/apps/tag"
	"hezihua/test/tools"
	"testing"
)

var (
	impl tag.Service
	ctx = context.Background()
)



func TestCreateTag(t *testing.T) {
	req := tag.NewCreateTagRequest("分类", "go", 1)
	ins, err := impl.CreateTag(ctx, req)	
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryTag(t *testing.T) {
	req := tag.NewQueryTagRequest()
	// req.Add(2)
	ins, err := impl.QueryTag(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}




func init() {
	tools.DevelopmentSet()
	impl = apps.GetInternalApp(tag.AppName).(tag.Service)
}