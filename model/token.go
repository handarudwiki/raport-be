package model

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AtExpires    int    `json:"at_expires"`
	RtExpires    int    `json:"rt_expires"`
}
