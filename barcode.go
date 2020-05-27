package sticker

import (
	"errors"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

// Barcode 条形码
type Barcode struct {
	Code string
	PtX  int
	PtY  int
}

// NewBarcode 条形码图片
func NewBarcode() *Barcode {
	return &Barcode{PtX: 0, PtY: 20}
}

// CreateBarcode 创建
func CreateBarcode(filePath, code string) (string, error) {
	if code == "" {
		return "", errors.New("code 不能为空")
		// code = "300684963555533158"
	}

	var fileFullPath string
	exists, err := PathExists(filePath)
	if err != nil {
		return "", err
	}
	if !exists {
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	// 创建一个code128编码的 BarcodeIntCS
	cs, _ := code128.Encode(code)
	// 创建一个要输出数据的文件
	fileFullPath = filePath + code + ".jpg"
	file, err := os.Create(fileFullPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 设置图片像素大小
	qrCode, err := barcode.Scale(cs, 270, 70)
	if err != nil {
		return "", err
	}
	// 将code128的条形码编码为png图片
	png.Encode(file, qrCode)

	return fileFullPath, nil
}

// PathExists  是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
