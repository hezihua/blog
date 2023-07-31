package tag

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"hezihua/validator"
	"time"
)


func NewTagSet() *TagSet {
	return &TagSet{
		Items: []*Tag{},
	}
}

func (s *TagSet) String() string {
	dj, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(dj)
}

type TagSet struct {
	Total int64 `json:"total"`
	Items []*Tag `json:"items"`
}

func NewTag (req *CreateTagRequest) (*Tag, error) {

	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	if err:= req.Validate(); err != nil {
		return nil, err
	}
	h := fnv.New32a()
	_, err := h.Write([]byte(fmt.Sprintf("%s.%s.%d", req.Key, req.Value, req.BlogId)))
	if err != nil {
		panic(err)
	}

	
	return &Tag{
		Id: base64.StdEncoding.EncodeToString(h.Sum(nil)),
		CreateTagRequest: req,
		CreatedAt: time.Now().Unix(),
	}, nil
}
type Tag struct {
	*CreateTagRequest
	CreatedAt int64 `json:"created_at"`
	Id       string   `json:"id"`
}

func (req *CreateTagRequest) Validate() error {
	return validator.V().Struct(req)
}