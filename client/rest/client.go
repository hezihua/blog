package rest

import (
	"github.com/infraboard/mcube/client/rest"
)

func NewClient(conf *Config) *Client {
	
	c := rest.NewRESTClient()
	c.SetBaseURL(conf.Url)
	c.SetBasicAuth(conf.Username, conf.Password)
	return &Client{c: c}
}

type Client struct {
	c *rest.RESTClient
}