package model

// WechatMiniUser WechatMiniUser
type WechatMiniUser struct {
	BaseControll
	UID        uint   `json:"uid"`
	OpenID     string `json:"openid" gorm:"unique"`
	Unionid    string `json:"unionid" gorm:""`
	SessionKey string `json:"session_key"`
}

// Pointer Pointer
func (w *WechatMiniUser) Pointer() interface{} {
	return &WechatMiniUser{}
}

// PointerList PointerList
func (w *WechatMiniUser) PointerList() interface{} {
	return &[]WechatMiniUser{}
}

// TableName TableName
func (w *WechatMiniUser) TableName() string { return TableName("wechat_mini_user") }
