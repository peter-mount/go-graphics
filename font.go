package graph

import (
	"image"
)

type Font interface {
	// DrawRune draws a rune onto an image.
	DrawRune(img Image, p image.Point, r rune) image.Point
	// Size returns the size of a rune
	Size(r rune) (int, int)
}
