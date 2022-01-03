package client

import (
	"fakeeyes_client/client/protos/request"
	"fmt"
	"testing"
)

var client *Client

var server = "http://127.0.0.1:8080/"

func init() {
	var err error
	client, err = NewClient(server)
	if err != nil {
		panic(err)
	}
}

func TestUserSignUp(t *testing.T) {

	name := "testuser"
	req := request.UserSignUp{
		Name: name,
	}
	resp, err := client.SignUp(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}