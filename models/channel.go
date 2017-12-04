package models

type Channel struct {
	ID           string   `json:"_id"`
	Name         string   `json:"name"`
	MessageCount int      `json:"msgs"`
	UserNames    []string `json:"usernames"`
	Default      bool     `json:"default"`

	User User `json:"u"`

	ReadOnly  bool   `json:"ro"`
	Timestamp string `json:"ts"`
	Type      string `json:"t"`
	UpdatedAt string `json:"_updatedAt"`
	SysMes    bool   `json:"sysMes"`
}

type ChannelSubscription struct {
	ID          string   `json:"_id"`
	Alert       bool     `json:"alert"`
	Name        string   `json:"name"`
	DisplayName string   `json:"fname"`
	Open        bool     `json:"open"`
	RoomID      string   `json:"rid"`
	Type        string   `json:"c"`
	User        User     `json:"u"`
	Roles       []string `json:"roles"`
	Unread      float64      `json:"unread"`
}


type ChannelUserInfo struct {
	Avatar string `json:"avatar"`
	Extra map[string]interface{} `json:"extra"`
	Name string `json:"name"`
	Status string `json:"status"`
	TZ int32 `json:"tz"`
	Roles []string `json:"roles"`
	Phone []UserPhone `json:"phone"`
	Email string `json:"email"`
}