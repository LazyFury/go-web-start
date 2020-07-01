package model

const (
	// Sigle 单图
	Sigle int = 1
	// Multi 多图
	Multi int = 2
)

// AdGroup 广告位
type AdGroup struct {
	Type int    `json:"type" gorm:"comment:'1单图,2多图'"`
	Name string `json:"name"`
	BaseControll
}

// AdEvent banner事件
type AdEvent struct {
	Event string `json:"event" gorm:"not null;unique_index;comment:'banner事件,字符串，唯一'"`
	BaseControll
}

// Ad 广告位
type Ad struct {
	URL     string `json:"url"`
	EventID int    `json:"event_id"`
	Title   string `json:"title"`
	BaseControll
}
