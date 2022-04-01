package graph

import (
	"image"
	"image/color"
)

// Mapper is a mapping function to map one Color to a new one.
// This is similar to the Convert function in the image.Model interface.
type Mapper func(col color.Color) color.Color

func MapperOf(mappers ...Mapper) Mapper {
	var r Mapper
	for _, m := range mappers {
		r = r.Then(m)
	}
	return r
}

func (m Mapper) Then(b Mapper) Mapper {
	if m == nil {
		return b
	}
	if b == nil {
		return m
	}
	return func(col color.Color) color.Color {
		return b(m(col))
	}
}

// Do applies the Mapper against the source image, writing the result to the destination
// over the area defined by the supplied rectangle.
func (m Mapper) Do(src image.Image, dst Image, b image.Rectangle) {
	if m != nil {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				dst.Set(x, y, m(src.At(x, y)))
			}
		}
	}
}

// DoNew applies the Mapper against an image, returning a new mutable image with the result.
func (m Mapper) DoNew(src image.Image) Image {
	dst := NewRGBA(src)
	m.Do(src, dst, dst.Bounds())
	return dst
}

// DoOver applies the Mapper over the supplied mutable image, overwriting its previous state.
func (m Mapper) DoOver(src Image) {
	m.Do(src, src, src.Bounds())
}
