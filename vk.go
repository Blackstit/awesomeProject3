package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	vkServiceKey = "77fd337377fd337377fd337397778eb76d777fd77fd337328c670c1009d987f845989fb"
	version      = "5.122"
	reqURL       = "https://api.vk.com/method/wall.get?"
)

type wallResponse struct {
	Body body `json:"response"`
}

type body struct {
	Items []Items `json:"items"`
}

type Items struct {
	Text string `json:"text"`
}

func getPosts(groupId string) ([]Items, error) {
	u := url.Values{}
	u.Set("count", "3")
	u.Set("offset", "0")
	u.Set("filter", "owner")
	u.Set("owner_id", groupId)
	u.Set("access_token", vkServiceKey)
	u.Set("v", version)

	req := reqURL + u.Encode()
	resp, err := http.Get(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := new(wallResponse)
	if err := json.Unmarshal(b, response); err != nil {
		return nil, err
	}
	return response.Body.Items, nil

}
