package rest

import (
	"fmt"
	"bytes"
	"errors"
	"net/http"
	"net/url"

	"github.com/tntobias/Rocket.Chat.Go.SDK/models"
)

type logoutResponse struct {
	StatusResponse
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}

type logonResponse struct {
	StatusResponse
	Data struct {
		Token  string `json:"authToken"`
		UserID string `json:"userId"`
	} `json:"data"`
}

type userInfoResponse struct {
	Success bool `json:"success"`
	UserInfo models.UserInfo `json:"user"`
}

// Login a user. The Email and the Password are mandatory. The auth token of the user is stored in the Client instance.
//
// https://rocket.chat/docs/developer-guides/rest-api/authentication/login
func (c *Client) Login(credentials models.UserCredentials) error {
	data := url.Values{"user": {credentials.Email}, "password": {credentials.Password}}
	request, _ := http.NewRequest("POST", c.getUrl()+"/api/v1/login", bytes.NewBufferString(data.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := new(logonResponse)

	if err := c.doRequest(request, response); err != nil {
		return err
	}

	if response.Status == "success" {
		c.auth = &authInfo{id: response.Data.UserID, token: response.Data.Token}
		return nil
	} 
	return errors.New("Response status: " + response.Status)

}

// Logout a user. The function returns the response message of the server.
//
// https://rocket.chat/docs/developer-guides/rest-api/authentication/logout
func (c *Client) Logout() (string, error) {

	if c.auth == nil {
		return "Was not logged in", nil
	}

	request, _ := http.NewRequest("POST", c.getUrl()+"/api/v1/logout", nil)

	response := new(logoutResponse)

	if err := c.doRequest(request, response); err != nil {
		return "", err
	}

	if response.Status == "success" {
		return response.Data.Message, nil
	} 
	return "", errors.New("Response status: " + response.Status)

}

func (c *Client) UserInfoByName(username string) (*models.UserInfo, error) {
	url := fmt.Sprintf("%s/api/v1/users.info?username=%s", c.getUrl(), username)
	request, _ := http.NewRequest("GET", url, nil)
	response := new(userInfoResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.UserInfo, nil
}
