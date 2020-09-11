package model

// Null Null
type Null struct {
	BaseControll
}

var _ Model = &Null{}

// Pointer Pointer
func (n *Null) Pointer() interface{} { return nil }

// PointerList PointerList
func (n *Null) PointerList() interface{} { return nil }

// TableName TableName
func (n *Null) TableName() string { return "null" }
