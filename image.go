// Package graph is a suite of utilities for performing various operation's on an Image
package graph

import (
	"image"
	"image/color"
	"image/draw"
)

// Image is an image.Image with a Set method to change a single pixel.
type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}

// NewRGBA creates a new mutable image with the same dimensions of another image
func NewRGBA(img image.Image) Image {
	return image.NewRGBA(img.Bounds())
}

// DuplicateImage creates a new copy of an image which is also mutable
func DuplicateImage(img image.Image) Image {
	dst := NewRGBA(img)
	draw.Draw(dst, img.Bounds(), img, image.Point{}, draw.Src)
	return dst
}

type wrapper struct {
	img image.Image
}

func (w *wrapper) ColorModel() color.Model     { return w.img.ColorModel() }
func (w *wrapper) Bounds() image.Rectangle     { return w.Bounds() }
func (w *wrapper) At(x, y int) color.Color     { return w.At(x, y) }
func (w *wrapper) Set(x, y int, c color.Color) {}

func Imutable(img image.Image) Image {
	return &wrapper{img: img}
}
