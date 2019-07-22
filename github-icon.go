package main

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"os"
	"strconv"
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
	HashKey    uint64
	ColorModel color.Color
}

func (avatar *Avatar) genHash() {
	h5 := md5.New()
	io.WriteString(h5, avatar.Key)
	avatar.HashKey, _ = strconv.ParseUint(fmt.Sprintf("%x", h5.Sum(nil)), 16, 64)
	fmt.Println(avatar.HashKey)
}

// https://golang.org/pkg/image/color/
func (avatar *Avatar) color() {
	//rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	avatar.ColorModel = color.RGBA{uint8(avatar.HashKey >> 4 & 0xff),
		uint8(avatar.HashKey >> 55 & 0xff),
		uint8(avatar.HashKey >> 58 & 0xff),
		uint8(avatar.HashKey >> 57 & 0xff)}
	fmt.Println(avatar.ColorModel)
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
	// 生成本身颜色
	avatar.genHash()
	avatar.color()
	avatar.HashKey >>= 16
	// 一共是25个格子
	x, y := 0, 0
	for walk := 0; walk < Size*Size; walk++ {
		if avatar.HashKey&1 == 0 {
			walk, _y := empty+x*Gird, empty+y*Gird
			avatar.addColor(walk, _y)
			// 对面也要画上
			walk = width - empty - 10 - walk
			avatar.addColor(walk, _y)
		}
		avatar.HashKey >>= 3
		y += 1
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
