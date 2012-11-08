// Package go-gcm can push a message to android phone through [Google Cloud Messaging(GCM)](http://developer.android.com/guide/google/gcm/index.html)
package gcm

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var GCMSendApi = "https://android.googleapis.com/gcm/send"

type Client struct {
	key  string
	http *http.Client
}

// Create a client with key. Grab the key from [Google APIs Console](https://code.google.com/apis/console)
func New(key string) *Client {
	return &Client{
		key:  key,
		http: new(http.Client),
	}
}

func (c *Client) Send(message *Message) (*Response, error) {
	j, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", GCMSendApi, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("key=%s", c.key))
	request.Header.Add("Content-Type", "application/json")

	resp, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}

	return responseReply(resp)
}

func httpClientWithoutSecureVerify() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	return &http.Client{Transport: tr}
}

func responseReply(resp *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s %s", resp.Status, string(body))
	}

	ret := new(Response)
	err = json.Unmarshal(body, ret)
	return ret, err
}
