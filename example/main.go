package main

import (
	"fmt"

	"github.com/jerryxjm/sticker"
)

func main() {
	s := sticker.New()
	s.Barcode.Code = "300684963555533158" //"300684963555533158"
	s.SavePath = s.StartPath() + "sticker/"
	s.SaveName = "A1204"
	// s.Size.X = 500
	// s.Size.Y = 600
	s.Barcode.PtX = 20

	font1 := sticker.NewFont()
	font1.PtX = 20
	font1.PtY = 100
	font1.LineTexts = []string{"300684963555533158"}
	font1.Size = 10
	s.Fonts = append(s.Fonts, font1)

	font2 := sticker.NewFont()
	font2.PtY = 120
	font2.LineTexts = []string{
		"CWSQ\\N档\\AS",
		"17241700\\200525-7-1\\1-1",
		"790#",
		"L\\绿色",
		"退货时请保证此标签完整\\韵达"}

	s.Fonts = append(s.Fonts, font2)

	err := s.Generate()
	if err != nil {
		fmt.Println(err)
	}
}
