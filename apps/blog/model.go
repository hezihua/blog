package blog

import (
	"encoding/json"
	"hezihua/validator"
	"time"
)

func NewBlog(req *CreateBlogRequest) *Blog {
	now := time.Now().Unix()
	return &Blog{
		Id: 		 0,
		CreatedAt: now,
		UpdatedAt: now,
		CreateBlogRequest: req,
	}
}

type Blog struct {
	Id       int64 `json:"id"`
	CreatedAt int64 `json:"created_at"`
	PulishedAt int64 `json:"pulished_at"`
	UpdatedAt int64 `json:"updated_at"`
	*CreateBlogRequest
}

func (b *Blog) String() string {
	dj, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return string(dj)
}

func (req *CreateBlogRequest) Validate() error {
	
	return validator.V().Struct(req)
}