package filter

import (
	"github.com/peter-mount/go-graphics/histogram"
	"github.com/peter-mount/go-graphics/util"
	"image/color"
)

// DeltaRGB contains the differences to apply to an RGBA colour in an AdjustImageFilter
type DeltaRGB struct {
	R, G, B int
}

func (d DeltaRGB) Apply(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	return util.RGBA(
		util.LimitU32(int(r)+d.R),
		util.LimitU32(int(g)+d.G),
		util.LimitU32(int(b)+d.B),
	)
}

// Filter creates a graph.Filter which will apply the DeltaRGB against an Image
func (d DeltaRGB) Filter(_ int, _ int, c color.Color) (color.Color, error) {
	return d.Apply(c), nil
}

func DeltaRGBFromHistogram(h *histogram.Histogram) DeltaRGB {
	var r, g, b uint32
	var rx, gx, bx uint8
	for x := 0; x < 256; x++ {
		r, rx = maxX(r, h.Red[x], rx, uint8(x))
		g, gx = maxX(g, h.Green[x], gx, uint8(x))
		b, bx = maxX(b, h.Blue[x], bx, uint8(x))
	}

	var c uint8
	if r > g && r > b {
		c = rx
	} else if g > r && g > b {
		c = gx
	} else if b > r && b > g {
		c = bx
	}

	return DeltaRGB{
		R: (int(c) - int(rx)) * 256,
		G: (int(c) - int(gx)) * 256,
		B: (int(c) - int(bx)) * 256,
	}
}

// maxX returns (a,ax) if a > b. If not then (b,bx) is returned
func maxX(a, b uint32, ax, bx uint8) (uint32, uint8) {
	if a > b {
		return a, ax
	}
	return b, bx
}
