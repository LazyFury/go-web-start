package wechat

const (
	//LoginURL 拼接微信登陆请求
	LoginURL string = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	//RedirectURL 跳转微信登陆授权页
	RedirectURL string = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect"

	//UserInfoURL 微信用户信息
	UserInfoURL string = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	// UserInfoCgiURL 用户信息
	UserInfoCgiURL string = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"

	//jsAPITicketURL jsapi_ticket授权请求
	jsAPITicketURL string = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"

	//accessTokenURL AccessToken 授权请求
	accessTokenURL string = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)
