package rocketchatgo

import (
	"github.com/tntobias/Rocket.Chat.Go.SDK/models"
	"github.com/tntobias/Rocket.Chat.Go.SDK/rest"
)

func GetAuthenticatedClient(host, port, username, email, password string, tls, debug bool) *rest.Client, err {
	c :=  rest.NewClient(host, port, tls, debug)
	cred := models.UserCredentials{
		Name: username,
		Email: email,
		Password: password,
	}

	if err := c.Login(cred); err != nil {
		return nil, err
	}

	return &c, nil

}