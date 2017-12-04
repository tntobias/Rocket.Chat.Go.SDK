package models


type User struct {
	ID       string `json:"_id"`
	UserName string `json:"username"`
}

type UserInfo struct {
	ID               string                 `json:"_id,omitempty"`
	CreatedAt        string                 `json:"createdAt,omitempty"`
	Emails           []UserEmail            `json:"emails,omitempty"`
	Phone			 []UserPhone  			`json:"phone" default:"{}"`
	Type             string                 `json:"type,omitempty"`
	Status           string                 `json:"status,omitempty"`
	Active           bool                   `json:"active,omitempty"`
	Roles            []string               `json:"roles,omitempty"`
	Extra	 map[string]interface{}          `json:"customFields,omitempty"`
	Name             string                 `json:"name,omitempty"`
	LastLogin        string                 `json:"lastLogin,omitempty"`
	StatusConnection string                 `json:"statusConnection,omitempty"`
	UTCOffset        int32                  `json:"utcOffset,omitempty"`
	Username         string                 `json:"username,omitempty"`
}

type UserEmail struct {
	Address string `json:"address,omitempty"`
	Verified bool `json:"verified,omitempty"`
}

type UserPhone struct {
	PhoneNumber string `json:"phoneNumber"`
}

func (x *UserInfo) SetDefaults() {
	np := make([]UserPhone, 1)
	np[0].PhoneNumber = "--"
	x.Phone = np
}
