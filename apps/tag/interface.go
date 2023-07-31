package tag

import "context"

var (
	AppName = "tags"
)
type Service interface {
	CreateTag(context.Context, *CreateTagRequest) (*Tag, error)
	QueryTag(context.Context, *QueryTagRequest) (*TagSet, error)
	RemoveTag(context.Context, *RemoveTagRequest) error
}

func NewCreateTagRequest(key, value string, blogId int) *CreateTagRequest {
	return &CreateTagRequest{
		Key: key,
		Value: value,
		BlogId: blogId,
	}
}

func NewDefaultCreateTagRequest() *CreateTagRequest {
	return &CreateTagRequest{}
}

type CreateTagRequest struct {
	Key      string `gorm:"column:t_key" json:"key" validate:"required"`
	Value    string `gorm:"column:t_value" json:"value" validate:"required"`
	BlogId   int `gorm:"column:blog_id" json:"blog_id" validate:"required"`
}

func NewQueryTagRequest() *QueryTagRequest {
	return &QueryTagRequest{
		BlogIds: []int{},
	}
}

func (req *QueryTagRequest) Add(ids ...int) {
	req.BlogIds = append(req.BlogIds, ids...)
}

type QueryTagRequest struct {
	BlogIds []int `json:"blog_id"`
}
type RemoveTagRequest struct {
	Id string `json:"id"`
}
