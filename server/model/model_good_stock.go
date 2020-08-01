package model

// GoodStock 库存
type GoodStock struct {
	BaseControll
}

// PointerList PointerList
func (g *GoodStock) PointerList() interface{} {
	return &[]GoodStock{}
}

// Pointer Pointer
func (g *GoodStock) Pointer() interface{} {
	return &GoodStock{}
}

// TableName TableName
func (g *GoodStock) TableName() string {
	return TableName("good_stocks")
}

// IsPublic 个人数据
func (g *GoodStock) IsPublic() bool { return false }
