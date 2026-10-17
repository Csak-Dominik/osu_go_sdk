package authentication

type Session interface {
	GetAccessToken() *string
}

type SessionService struct {
	Session Session
}

type OAuthAuthorizationRequest struct {
	// The Client ID you received when you [registered].
	//
	// [registered]: https://osu.ppy.sh/home/account/edit#new-oauth-application
	ClientID int64 `json:"client_id"`
	// Optional. The URL in your application where users will be sent after authorization. This must match the registered Application Callback URL exactly.
	RedirectURI *string `json:"redirect_uri"`
	// Optional. A list of [Scope].
	Scopes []string `json:"scopes"`
	// Optional. Data that will be returned when a temporary code is issued. It can be used to provide a token for protecting against cross-site request forgery attacks.
	State string `json:"state"`
}
