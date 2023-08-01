package rest

import (
	"context"
	"fmt"
	"hezihua/apps/user"
	"net/http"
	"time"

	"github.com/infraboard/mcube/client/rest"
)

func NewClient(conf *Config) *Client {
	
	c := rest.NewRESTClient()
	c.SetBaseURL(conf.Url)
	c.SetBasicAuth(conf.Username, conf.Password)
	c.SetCookie(&http.Cookie{Name: "username", Value: conf.Username})
	c.SetBasicAuth(conf.Username, conf.Password)
	return &Client{c: c}
}

type Response struct { 
	Code int `json:"code"`
	Message *user.Session `json:"message"`
}


func (c *Client) Session () string {
	return fmt.Sprintf("username: %s, password: %s", 
	c.session.Username, c.session.Id)
}
func (c *Client) Login() error{
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp := &Response{Message: &user.Session{}}
	err := c.c.Post("/user/login").Do(ctx).Into(resp)
	if err != nil {
		return err
	}
	c.session = resp.Message
	c.c.SetCookie(&http.Cookie{
		Name: "username", 
		Value: c.session.Username,
	}, &http.Cookie{
		Name: "session",
		Value: c.session.Id,
	})
	fmt.Println(*resp, ctx)
	return nil
}

type Client struct {
	c *rest.RESTClient
	session *user.Session
}