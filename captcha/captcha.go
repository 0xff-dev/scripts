package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	char   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	width  = 200
	height = 100
)

func randColor() color.RGBA {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return color.RGBA{uint8(rnd.Intn(256)), uint8(rnd.Intn(256)),
		uint8(rnd.Intn(256)), uint8(rnd.Intn(256))}
}

func randChar() byte {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return char[rnd.Intn(52)]
}

// GenerateImg 生成验证码
func GenerateImg() {
	upLeft, lowRight := image.Point{0, 0}, image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, randColor())
		}
	}
	fontFile, err := os.Open("./Krungthep.ttf")
	if err != nil {
		log.Fatal("Open font file error.")
		return
	}
	defer fontFile.Close()
	fontByte, err := ioutil.ReadAll(fontFile)
	if err != nil {
		log.Fatal("Read font file error")
		return
	}
	ft, err := truetype.Parse(fontByte)
	d := &font.Drawer{
		Dst: img,
		Src: image.Black,
		Face: truetype.NewFace(ft, &truetype.Options{
			Size:    22,
			DPI:     100,
			Hinting: font.HintingNone,
		}),
	}
	for x := 0; x < width; x += 50 {
		d.Dot = fixed.P(x+15, height/2+10)
		d.DrawString(string(randChar()))
	}
	if res, err := os.Create("res.jpg"); err == nil {
		defer res.Close()
		err = jpeg.Encode(res, img, &jpeg.Options{90})
	}
}

func main() {
	GenerateImg()
}
