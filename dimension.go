package graph

import "image"

type Dimension struct {
	Width, Height int
}

func NewDimension(w, h int) Dimension {
	return Dimension{Width: w, Height: h}
}

func DimensionFromRect(rect image.Rectangle) Dimension {
	return NewDimension(rect.Max.X-rect.Min.X+1, rect.Max.Y-rect.Min.Y+1)
}

func (d Dimension) Area() int {
	return d.Width * d.Height
}

func (d Dimension) Rect(x, y int) image.Rectangle {
	return image.Rect(x, y, x+d.Width-1, y+d.Height-1)
}
