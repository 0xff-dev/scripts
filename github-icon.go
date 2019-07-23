package main

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"math/rand"
	"os"
	"time"
)

// 开始的位置是{20, 20}, 结束的位置{170, 170}
// 格子的宽度是30, 边界的宽度是20空出20的空间
const (
	width   = 190
	height  = 190
	Gird    = 30
	empty   = 20
	Size    = 5
	bgColor = "#fff"
)

type Avatar struct {
	Img        *image.RGBA
	Key        string
	HashKey    []byte
	ColorModel color.Color
}

func (avatar *Avatar) genHash() {
	h5 := md5.New()
	io.WriteString(h5, avatar.Key)
	avatar.HashKey = h5.Sum(nil)
}

// https://golang.org/pkg/image/color/
func (avatar *Avatar) color() {
	timeLayout, timeFmt := "2006-01-02 15:04:05", "1999-10-10 10:10:10"
	loc, _ := time.LoadLocation("Local") //获取时区
	_time, _ := time.ParseInLocation(timeLayout, timeFmt, loc)
	rnd := rand.New(rand.NewSource(_time.Unix()))
	avatar.ColorModel = color.RGBA{avatar.HashKey[rnd.Intn(16)] & 0xff,
		avatar.HashKey[rnd.Intn(16)] & 0xff,
		avatar.HashKey[rnd.Intn(16)] & 0xff,
		avatar.HashKey[rnd.Intn(16)] & 0xff}
}

func (avatar *Avatar) addColor(x, y int) {
	for _x := x; _x < x+Gird; _x++ {
		for _y := y; _y < y+Gird; _y++ {
			avatar.Img.Set(_x, _y, avatar.ColorModel)
		}
	}
}

func (avatar *Avatar) Draw(path string) {
	avatar.Img = image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	avatar.genHash()
	avatar.color()
	x, y, key := 0, 0, 0
	for walk := 0; walk < Size*Size; walk++ {
		if avatar.HashKey[key]&1 == 0 {
			walk, _y := empty+x*Gird, empty+y*Gird
			avatar.addColor(walk, _y)
			walk = width - empty - 10 - walk
			avatar.addColor(walk, _y)
		}
		y, key = y+1, (key+1)%16
		if y == Size {
			y = 0
			x += 1
		}
	}
	if res, err := os.Create(path + fmt.Sprintf("/%s.jpg", avatar.Key)); err == nil {
		defer res.Close()
		err = jpeg.Encode(res, avatar.Img, &jpeg.Options{90})
	}
}

func main() {
	av := &Avatar{Key: "150405211"}
	av.Draw("./")
}
