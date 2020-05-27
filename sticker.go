package sticker

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

// Sticker 贴纸
type Sticker struct {
	startPath string
	Barcode   string
	SavePath  string
	SaveName  string
	LineTexts []string
	Size      *Size
	Font      *Font
}

// New 贴纸
func New() *Sticker {
	s := &Sticker{Size: NewSize(), Font: NewFont()}
	var err error
	s.startPath, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	s.startPath = s.startPath + "/"
	return s
}

// Generate 生成
func (s *Sticker) Generate() error {
	// 生成条形码
	barcodeFileFullPath, err := CreateBarcode(s.startPath+"barcode/", s.Barcode)
	if err != nil {
		return err
	}

	exists, err := PathExists(s.SavePath)
	if err != nil {
		return err
	}
	if !exists {
		err = os.MkdirAll(s.SavePath, os.ModePerm)
		if err != nil {
			return err
		}
		err = os.Chmod(s.SavePath, 0777)
		if err != nil {
			return err
		}
	}
	if !strings.HasSuffix(s.SavePath, "/") {
		s.SavePath = s.SavePath + "/"
	}

	file, err := os.Create(s.SavePath + s.SaveName + ".jpg")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	file1, err := os.Open(barcodeFileFullPath)
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	img, _ := png.Decode(file1)
	//尺寸
	// img = resize.Resize(314, 314, img, resize.Lanczos3)

	// jpg := image.NewRGBA(image.Rect(0, 0, 827, 1169))
	jpg := image.NewRGBA(image.Rect(0, 0, s.Size.X, s.Size.Y))

	fontRender(jpg, s.Font, s.LineTexts)

	draw.Draw(jpg, img.Bounds().Add(image.Pt(0, 20)), img, img.Bounds().Min, draw.Src) //截取图片的一部分
	// draw.Draw(jpg, img.Bounds().Add(image.Pt(435, 150)), img, img.Bounds().Min, draw.Src) //截取图片的一部分
	// draw.Draw(jpg, img.Bounds().Add(image.Pt(60, 610)), img, img.Bounds().Min, draw.Src)  //截取图片的一部分
	// draw.Draw(jpg, img.Bounds().Add(image.Pt(435, 610)), img, img.Bounds().Min, draw.Src) //截取图片的一部分

	png.Encode(file, jpg)
	return nil
}

// StartPath 启动路劲
func (s *Sticker) StartPath() string {
	return s.startPath
}

func fontRender(jpg *image.RGBA, stickerFont *Font, lineTexts []string) {
	fontBytes, err := ioutil.ReadFile(stickerFont.FilePath)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	fg, bg := image.Black, image.White
	//ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	//if *wonb {
	//	fg, bg = image.White, image.Black
	//	ruler = color.RGBA{0x22, 0x22, 0x22, 0xff}
	//}
	draw.Draw(jpg, jpg.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(stickerFont.DPI)
	c.SetFont(f)
	c.SetFontSize(stickerFont.Size)
	c.SetClip(jpg.Bounds())
	c.SetDst(jpg)
	c.SetSrc(fg)

	switch stickerFont.Hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	//Draw the guidelines.
	//for i := 0; i < 200; i++ {
	//	jpg.Set(10, 10+i, ruler)
	//	jpg.Set(10+i, 10, ruler)
	//}

	// Draw the text.
	pt := freetype.Pt(20, 100+int(c.PointToFixed(stickerFont.Size)>>6))
	for _, s := range lineTexts {
		_, err = c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(stickerFont.Size * stickerFont.Spacing)
	}
}
