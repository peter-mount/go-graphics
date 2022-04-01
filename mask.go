package graph

import (
	"image"
	"image/draw"
)

// MaskImage applies an image mask against an Image.
// Normally the mask is a 1 bit image where black pixels are drawn over the top of
// the source image removing areas of the image you do not want to be shown.
func MaskImage(src image.Image, mask image.Image) Image {
	// The new image bounds - this is the smallest rectangle contained by both
	b := mask.Bounds().Union(src.Bounds())

	//dst := image.NewRGBA(b)
	//draw.Draw(dst, b, src, image.Point{}, draw.Src)
	dst := DuplicateImage(src)
	draw.Draw(dst, b, mask, image.Point{}, draw.Over)
	draw.DrawMask(dst, b, mask, image.Point{}, src, image.Point{}, draw.Over)
	return dst
}
