package wechat

// UserInfo 用户信息
type UserInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Nickname     string `json:"nickname"`
	Sex          int    `json:"sex"`
	Headimgurl   string `json:"headimgurl"`
	Province     string `json:"province"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Unionid      string `json:"unionid"`
	Subscribe    int    `json:"subscribe"`
}
