package client

import (
	"encoding/json"
	"fakeeyes_client/client/protos/response"

	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Client struct {
	Server string
	Token  string
	url    *url.URL
}

var APIPrefix = "/api/v1"

func NewClient(server string) (*Client, error) {
	client := Client{
		Server: server,
	}
	url, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	client.url = url

	return &client, nil
}

//
func (c *Client) httpproxy(api string, req interface{}, resp interface{}) error {
	var err error

	relurl := path.Join(APIPrefix, api)
	u, err := url.Parse(relurl)
	if err != nil {
		return err
	}
	queryURL := c.url.ResolveReference(u).String()

	var reqstr string
	if req == nil {
		reqstr = ""
	} else {

		reqbody, err := json.Marshal(req)
		if err != nil {
			return err
		}
		reqstr = string(reqbody)
	}
	httpreq, err := http.NewRequest(http.MethodPost, queryURL, strings.NewReader(reqstr))
	if err != nil {
		return err
	}
	// sb := c.sign(reqstr)
	// httpreq.Header.Add("Timestamp", fmt.Sprintf("%d", sb.Timestamp))
	// httpreq.Header.Add("Signature", sb.Sign)
	// httpreq.Header.Add("Accesskey", sb.Accesskey)
	httpclt := http.Client{}
	httpresp, err := httpclt.Do(httpreq)
	if err != nil {
		return err
	}
	defer httpresp.Body.Close()
	body, err := ioutil.ReadAll(httpresp.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(body))
	if httpresp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Code: %d , HTTP Response: %s ", httpresp.StatusCode, string(body))
	}
	rm := response.ReturnMessage{}

	err = json.Unmarshal(body, &rm)
	if err != nil {
		return err
	}
	if !rm.Success {
		err = fmt.Errorf("ErrorCode: %s  ErrorMessage: %s", rm.ErrorCode, rm.ErrorMessage)
		return err
	}
	err = json.Unmarshal(body, resp)
	return err
}

//
func (c *Client) NewUser(token string) *User {

	user := User{
		client: c,
		token:  token,
	}
	return &user
}
