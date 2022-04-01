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
