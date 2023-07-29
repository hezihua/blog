package blog

import "encoding/json"

type Blog struct {
	Id       int64 `json:"id"`
	CreateAt int64 `json:"create_at"`
	PulishAt int64 `json:"pulish_at"`
	UpdateAt int64 `json:"update_at"`
	*CreateBlogRequest
}

func (b *Blog) String() string {
	dj, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return string(dj)
}