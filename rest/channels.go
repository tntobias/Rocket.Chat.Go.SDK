package rest

import (
	"bytes"
	"fmt"
	"net/http"
	"log"

	"github.com/tntobias/Rocket.Chat.Go.SDK/models"
	"github.com/creasty/defaults"
)

type channelsResponse struct {
	Success  bool             `json:"success"`
	Channels []models.Channel `json:"channels"`
}

type channelResponse struct {
	Success bool           `json:"success"`
	Channel models.Channel `json:"channel"`
}

// Returns all channels that can be seen by the logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list
func (c *Client) GetPublicChannels() ([]models.Channel, error) {
	request, _ := http.NewRequest("GET", c.getUrl()+"/api/v1/channels.list", nil)
	response := new(channelsResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

// Returns all channels that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list-joined
func (c *Client) GetJoinedChannels() ([]models.Channel, error) {
	request, _ := http.NewRequest("GET", c.getUrl()+"/api/v1/channels.list.joined", nil)
	response := new(channelsResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

// Joins a channel. The id of the channel has to be not nil.
//
// This function is not supported by the current Client.Chat release version 0.48.2.
func (c *Client) JoinChannel(channel *models.Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s" }`, channel.ID)
	request, _ := http.NewRequest("POST", c.getUrl()+"/api/v1/channels.join", bytes.NewBufferString(body))
	return c.doRequest(request, new(StatusResponse))
}

// Leaves a channel. The id of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/leave
func (c *Client) LeaveChannel(channel *models.Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s"}`, channel.ID)
	request, _ := http.NewRequest("POST", c.getUrl()+"/api/v1/channels.leave", bytes.NewBufferString(body))
	return c.doRequest(request, new(StatusResponse))
}

// Get information about a channel. That might be useful to update the usernames.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/info
func (c *Client) GetChannelInfo(channel *models.Channel) (*models.Channel, error) {
	var url = fmt.Sprintf("%s/api/v1/channels.info?roomId=%s", c.getUrl(), channel.ID)
	request, _ := http.NewRequest("GET", url, nil)
	response := new(channelResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Channel, nil
}

func (c *Client) GetChannelInfoByName(channel string) (*models.Channel, error) {
	var url = fmt.Sprintf("%s/api/v1/channels.info?roomName=%s", c.getUrl(), channel)
	request, _ := http.NewRequest("GET", url, nil)
	response := new(channelResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Channel, nil
}

func (c *Client) GetChannelUserInfo(channel string) (*[]models.ChannelUserInfo, error) {

	ch, err := c.GetChannelInfoByName(channel)
	if err != nil {
		return nil, err
	}
	log.Println(len(ch.UserNames))
	res := make([]models.ChannelUserInfo, 0)
	for _, u := range ch.UserNames {
		uInfo, err := c.UserInfoByName(u)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling %s: %s", u, err.Error())
		}
		if err := defaults.Set(uInfo); err != nil {
			return nil, err
		}
		res = append(res, models.ChannelUserInfo{
			Avatar: fmt.Sprintf("%s/avatar/%s", c.getUrl(), u),
			Extra: uInfo.Extra,
			Name: uInfo.Name,
			Status: uInfo.Status,
			TZ: uInfo.UTCOffset,
			Roles: uInfo.Roles,
			Phone: uInfo.Phone,
			Email: uInfo.Emails[0].Address,
		})
	}

	return &res, nil
}
