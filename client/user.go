package client

import (
	"fakeeyes_client/client/protos/request"
	"fakeeyes_client/client/protos/response"
)

type User struct {
	client *Client
	token  string
}

// 用户登陆
func (c *Client) SignIn(req request.UserSignIn) (user *User, err error) {

	// var err error
	apiname := "/User/SignIn"

	type rs struct {
		Data response.UserLogin
	}
	resp := &rs{}

	err = c.httpproxy(apiname, req, &resp)
	if err != nil {
		return
	}
	user = c.NewUser(resp.Data.Token)
	return
}

// 用户注册
func (c *Client) SignUp(req request.UserSignUp) (user *User, err error) {

	apiname := "/User/SignUp"

	type rs struct {
		Data response.UserLogin
	}
	resp := &rs{}

	err = c.httpproxy(apiname, req, &resp)
	if err != nil {
		return
	}
	user = c.NewUser(resp.Data.Token)
	return

}
