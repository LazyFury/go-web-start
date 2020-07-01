package model

// FeedBack 客户意见反馈
type FeedBack struct {
	BaseControll
	Reason  string `json:"reason"`
	Content string `json:"content" gorm:"type:text"`
}
