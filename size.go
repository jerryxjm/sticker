package sticker

// Size 尺寸
type Size struct {
	X int
	Y int
}

// NewSize 图片大小
func NewSize() *Size {
	return &Size{X: 340, Y: 257}
}
