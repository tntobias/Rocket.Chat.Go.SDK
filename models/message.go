package models

type Message struct {
	ID        string `json:"_id"`
	ChannelID string `json:"rid"`
	Text      string `json:"msg"`
	Timestamp string `json:"ts"`
	User      User   `json:"u"`
}
