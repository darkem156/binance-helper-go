package utils

import (
	"encoding/json"
)

func (client *Client) SendPublicRequest(request string, params map[string]string) string {
	response, _ := json.Marshal(params)
	return string(response)
}

func (client *Client) SendSignedRequest(request string, params map[string]string) string {
	return ""
}
