package model

// GoodSku 商品规格
type GoodSku struct {
	BaseControll
}

// PointerList PointerList
func (g *GoodSku) PointerList() interface{} {
	return &[]GoodSku{}
}

// Pointer Pointer
func (g *GoodSku) Pointer() interface{} {
	return &GoodSku{}
}

// TableName TableName
func (g *GoodSku) TableName() string {
	return TableName("good_skus")
}

// IsPublic 个人数据
func (g *GoodSku) IsPublic() bool { return false }
