package osugosdk

import (
	"github.com/Csak-Dominik/osu_go_sdk/authentication"
)

type Client struct {
}

func NewClient() Client {
	return Client{}
}

func (c Client) NewSession(accessToken *string) *ApiSession {
	session := &ApiSession{AccessToken: accessToken}

	session.Authentication = &authentication.SessionService{
		Session: session,
	}

	return session
}
