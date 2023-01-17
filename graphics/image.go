package graphics

import (
	graph "github.com/peter-mount/go-graphics"
	"github.com/peter-mount/go-graphics/util"
	"image"
	"image/draw"
)

func (g *graphics) Image() graph.Image {
	return g.img
}

func (g *graphics) Bounds() image.Rectangle {
	return g.bounds
}

func (g *graphics) Top() int {
	return g.bounds.Min.Y
}

func (g *graphics) Left() int {
	return g.bounds.Min.X
}

func (g *graphics) Width() int {
	return g.bounds.Dx()
}

func (g *graphics) Height() int {
	return g.bounds.Dy()
}

func (g *graphics) Crop(bounds image.Rectangle) graph.Graphics {
	dstImage := image.NewRGBA(bounds.Sub(bounds.Min))
	draw.Draw(dstImage, dstImage.Bounds(), g.img, bounds.Min, draw.Src)
	g.img = dstImage
	return g
}

func (g *graphics) Expand(top, left, bottom, right int) graph.Graphics {
	// Ensure parameters are >=0
	left = util.Max(left, 0)
	right = util.Max(right, 0)
	top = util.Max(top, 0)
	bottom = util.Max(bottom, 0)

	// Calculate new image size
	oldBounds := g.Bounds()
	newBounds := image.Rectangle{Min: oldBounds.Min, Max: image.Point{X: oldBounds.Max.X + left + right, Y: oldBounds.Max.Y + top + bottom}}

	// Rectangle to draw old image into new one
	topLeft := image.Point{X: left, Y: top}
	drawBounds := oldBounds.Add(topLeft)

	dstImage := image.NewRGBA(newBounds)
	draw.Draw(dstImage, drawBounds, g.img, image.Point{}, draw.Src)

	g.img = dstImage
	return g
}
