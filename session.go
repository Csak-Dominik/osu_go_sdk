package osugosdk

import "github.com/Csak-Dominik/osu_go_sdk/authentication"

type ApiSession struct {
	AccessToken *string

	// services
	Authentication *authentication.SessionService
}

func (s *ApiSession) GetAccessToken() *string {
	return s.AccessToken
}
