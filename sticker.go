package sticker

import (
	"image"
	"image/color"
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
	SavePath  string
	SaveName  string
	Size      *Size
	Fonts     []*Font
	Barcode   *Barcode
}

// New 贴纸
func New() *Sticker {
	s := &Sticker{Size: NewSize(), Fonts: []*Font{}, Barcode: NewBarcode()}
	var err error
	s.startPath, err = os.Getwd()
	if err != nil {
		return nil
	}
	s.startPath = s.startPath + "/"
	return s
}

// Generate 生成
func (s *Sticker) Generate() error {
	// 生成条形码
	barcodeFileFullPath, err := CreateBarcode(s.startPath+"barcode/", s.Barcode.Code, s.Barcode.Width, s.Barcode.Height)
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
		return err
	}
	defer file.Close()

	file1, err := os.Open(barcodeFileFullPath)
	if err != nil {
		return err
	}
	defer file1.Close()
	img, _ := png.Decode(file1)
	//尺寸
	// img = resize.Resize(314, 70, img, resize.Lanczos3)

	// jpg := image.NewRGBA(image.Rect(0, 0, 827, 1169))
	jpg := image.NewRGBA(image.Rect(0, 0, s.Size.X, s.Size.Y))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(jpg, jpg.Bounds(), &image.Uniform{white}, image.ZP, draw.Src) //画上白色背景

	for _, item := range s.Fonts {
		fontRender(jpg, item)
	}

	draw.Draw(jpg, img.Bounds().Add(image.Pt(s.Barcode.PtX, s.Barcode.PtY)), img, img.Bounds().Min, draw.Src) //截取图片的一部分
	// draw.Draw(jpg, img.Bounds().Add(image.Pt(435, 150)), img, img.Bounds().Min, draw.Src) //截取图片的一部分
	// draw.Draw(jpg, img.Bounds().Add(image.Pt(60, 610)), img, img.Bounds().Min, draw.Src)  //截取图片的一部分
	// draw.Draw(jpg, img.Bounds().Add(image.Pt(435, 610)), img, img.Bounds().Min, draw.Src) //截取图片的一部分

	png.Encode(file, jpg)
	os.Remove(barcodeFileFullPath)
	return nil
}

// StartPath 启动路劲
func (s *Sticker) StartPath() string {
	return s.startPath
}

func fontRender(jpg *image.RGBA, stickerFont *Font) {
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

	bounds := jpg.Bounds()

	draw.Draw(jpg, image.Rect(stickerFont.PtX, stickerFont.PtY, bounds.Max.X, bounds.Max.Y), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(stickerFont.DPI)
	c.SetFont(f)
	c.SetFontSize(stickerFont.Size)
	c.SetClip(image.Rect(stickerFont.PtX, stickerFont.PtY, bounds.Max.X, bounds.Max.Y))
	c.SetDst(jpg)
	c.SetSrc(fg)

	switch stickerFont.Hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	// Draw the text.
	pt := freetype.Pt(stickerFont.PtX, stickerFont.PtY+int(c.PointToFixed(stickerFont.Size)>>6))
	for _, s := range stickerFont.LineTexts {
		_, err = c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(stickerFont.Size * stickerFont.Spacing)
	}
}
