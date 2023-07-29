package blog

type Status int8

const (
	StatusDraft Status = iota
	StatusPublished
)

type UpdateMode string

const (
	PATCH = "patch"
	PUT   = "put"
)