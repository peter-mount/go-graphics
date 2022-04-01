package draw

import (
	"github.com/peter-mount/go-graphics"
	"image"
	"image/color"
)

// DrawRect draws an image.Rectangle
func DrawRect(img graph.Image, rect image.Rectangle, col color.Color) {
	for x := rect.Min.X; x < rect.Max.X; x++ {
		img.Set(x, rect.Min.Y, col)
		img.Set(x, rect.Max.Y, col)
	}
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		img.Set(rect.Min.X, y, col)
		img.Set(rect.Max.X, y, col)
	}
}

// FillRect draws a filled image.Rectangle
func FillRect(img graph.Image, rect image.Rectangle, col color.Color) {
	for y := rect.Min.Y; y <= rect.Max.Y; y++ {
		for x := rect.Min.X; x <= rect.Max.X; x++ {
			img.Set(x, y, col)
		}
	}
}
