package model

// 点赞类型声明
const (
	LIKE    int = 1
	UNLIKE  int = 2
	COMMENT int = 3
	ORDER   int = 4
	BUY     int = 5
	ADDCART int = 6
	SHARE   int = 7
)

// MessageType 操作类型
var MessageType = map[int]string{
	LIKE:    "点赞",
	UNLIKE:  "取消点赞",
	COMMENT: "评论",
	ORDER:   "下单",
	BUY:     "购买",
	ADDCART: "添加购物车",
	SHARE:   "分享",
}

// Message 客户意见反馈
// 暂定 如果存sql数据量太多，后期尝试redis之类的
type Message struct {
	BaseControll
	OrderID   uint `json:"orderId"`
	ArticleID uint `json:"articleId"`
	Action    int  `json:"action"`
	Remark    int  `json:"remark"`
}
