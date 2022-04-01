package histogram

import (
	"github.com/peter-mount/go-graphics"
	draw2 "github.com/peter-mount/go-graphics/draw"
	"github.com/peter-mount/go-graphics/graphics"
	"github.com/peter-mount/go-graphics/util"
	"image"
	"image/color"
	"image/draw"
)

func (h *Histogram) Plot(p image.Point, img draw.Image) *Histogram {
	graphics.New(img).Draw(h.Drawable(p))
	return h
}

func (h *Histogram) CompileDrawable(p image.Point) []graph.Drawable {
	minR, maxR := util.MinMax(h.Red)
	minG, maxG := util.MinMax(h.Green)
	minB, maxB := util.MinMax(h.Blue)
	min := util.MinRGB(minR, minG, minB)
	max := util.MaxRGB(maxR, maxG, maxB)
	fac := 256.0 / float64(max-min)

	return []graph.Drawable{h.polygon(p, fac, h.Red), h.polygon(p, fac, h.Green), h.polygon(p, fac, h.Blue)}
}

type drawable struct {
	p    image.Point
	d    []graph.Drawable
	rect image.Rectangle
}

func (h *Histogram) Drawable(p image.Point) graph.Drawable {
	return &drawable{
		p:    p,
		d:    h.CompileDrawable(p),
		rect: image.Rect(p.X, p.Y, p.X+257, p.Y+257),
	}
}

func (d *drawable) Draw(g graph.Graphics) {
	g.Background(color.Gray16{Y: 0x3333}).
		FillRectangle(d.rect).
		Foreground(color.White).
		DrawRectangle(d.rect).
		Foreground(color.RGBA{R: 255, A: 255}).Draw(d.d[0]).
		Foreground(color.RGBA{G: 255, A: 255}).Draw(d.d[1]).
		Foreground(color.RGBA{B: 255, A: 255}).Draw(d.d[2])
}

func (h *Histogram) polygon(p image.Point, fac float64, a []uint32) draw2.Polygon {
	poly := draw2.NewPolygon(false)
	//var ly int
	for x, v := range a {
		poly.Add(p.X+x-1, p.Y+256-int(float64(v)*fac))
		/*y := p.Y + 256 - int(float64(v)*fac)
		  if x > 0 {
		    graph.DrawLine(img, p.X+x-1, ly, p.X+x, y, c)
		  }
		  ly = y*/
	}
	return poly
}
