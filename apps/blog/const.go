package blog

import (
	"fmt"
	"strings"
)

func ParseStatusFromString(status string) (*Status, error) {
	for k, v := range StatusValueMap {
		if v == status {
			return &k, nil
		}
	}
	return nil, fmt.Errorf("no support %s", status)
}

type Status int8

func (s *Status) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = strings.Trim(str, `"`)
	for k, v := range StatusValueMap {
		if v == str {
			*s = k
			return nil
		}
	}
	return nil
}

func (s Status) MarshalJSON() ([]byte, error) {
	v := StatusValueMap[s]
	v = fmt.Sprintf(`"%s"`, v)
	return []byte(v), nil
}

const (
	StatusDraft Status = iota
	StatusPublished
)

var (
	StatusValueMap = map[Status]string{
		StatusDraft:     "draft",
		StatusPublished: "published",
	}
)

type UpdateMode string

const (
	PATCH UpdateMode = "patch"
	PUT UpdateMode  = "put"
)