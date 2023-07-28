package tag

import "context"

var (
	AppName = "tag"
)
type Service interface {
	CreateTag(context.Context, *CreateTagRequest) (*CreateTagResponse, error)
}

type CreateTagRequest struct {

}

type CreateTagResponse struct {

}