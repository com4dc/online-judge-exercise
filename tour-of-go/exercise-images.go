package main

import (
	"image"       // 追加
	"image/color" // 追加

	"golang.org/x/tour/pic"
)

type Image struct{}

// この場合はImageの中の値を書き換えてるわけではないため、ポインタ使わない
func (a Image) ColorModel() color.Model {
	return color.RGBAModel
}

// 指定されたインタフェースを
func (a Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 100)
}

// 満たしていくために実装していく
func (a Image) At(x, y int) color.Color {
	v := uint8((x ^ y) * 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
