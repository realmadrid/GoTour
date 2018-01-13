package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct {
	Width  int
	Height int
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(x, y int) color.Color {
	z := (uint8)(x+y) / 2
	return color.RGBA{z, z, 255, 255}
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
