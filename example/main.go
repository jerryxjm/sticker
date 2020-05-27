package main

import (
	"fmt"

	"github.com/jerryxjm/sticker"
)

func main() {
	s := sticker.New()
	s.Barcode.Code = "300684963555533158"
	s.SavePath = s.StartPath() + "sticker/"
	s.SaveName = "A1204"
	s.LineTexts = []string{
		"CWSQ\\N档\\AS",
		"17241700\\200525-7-1\\1-1",
		"790#",
		"L\\绿色",
		"请保证此标签完整\\韵达",
	}
	err := s.Generate()
	if err != nil {
		fmt.Println(err)
	}
}
