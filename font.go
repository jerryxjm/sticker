package sticker

// Font 字体信息
type Font struct {
	FilePath     string
	Size         float64
	Hinting      string
	Spacing      float64
	Whiteonblack bool
	DPI          float64
}

// NewFont 字体信息实例
func NewFont() *Font {
	font := &Font{}
	font.FilePath = "runtime/fonts/msyh.ttf"
	font.Size = 16
	font.Hinting = "none"     // none | full
	font.Spacing = 1.5        // line spacing (e.g. 2 means double spaced)
	font.Whiteonblack = false // white text on a black background
	font.DPI = 72
	return font
}
