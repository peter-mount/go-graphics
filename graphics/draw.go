package graphics

import (
	"github.com/peter-mount/go-graphics"
	"github.com/peter-mount/go-graphics/draw"
	"image"
	draw2 "image/draw"
)

func (g *graphics) Plot(x, y int) graph.Graphics {
	g.img.Set(x, y, g.foreground)
	return g
}
func (g *graphics) PlotPoint(p image.Point) graph.Graphics {
	return g.Plot(p.X, p.Y)
}

func (g *graphics) DrawBetweenPoints(p1, p2 image.Point) graph.Graphics {
	draw.DrawBetweenPoints(g.img, p1, p2, g.foreground)
	return g
}

func (g *graphics) DrawLine(x1, y1, x2, y2 int) graph.Graphics {
	draw.DrawLine(g.img, x1, y1, x2, y2, g.foreground)
	return g
}

func (g *graphics) Draw(d ...graph.Drawable) graph.Graphics {
	for _, e := range d {
		e.Draw(g)
	}
	return g
}

func (g *graphics) DrawRect(x1, y1, x2, y2 int) graph.Graphics {
	return g.DrawRectangle(image.Rect(x1, y1, x2, y2))
}

func (g *graphics) DrawRectangle(rect image.Rectangle) graph.Graphics {
	draw.DrawRect(g.img, rect, g.foreground)
	return g
}

func (g *graphics) FillRect(x1, y1, x2, y2 int) graph.Graphics {
	return g.FillRectangle(image.Rect(x1, y1, x2, y2))
}

func (g *graphics) FillRectangle(rect image.Rectangle) graph.Graphics {
	draw.FillRect(g.img, rect, g.background)
	return g
}

func (g *graphics) DrawImage(r image.Rectangle, src image.Image, sp image.Point, op draw2.Op) graph.Graphics {
	draw2.Draw(g.img, r, src, sp, op)
	return g
}

func (g *graphics) DrawMask(r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op draw2.Op) graph.Graphics {
	draw2.DrawMask(g.img, r, src, sp, mask, mp, op)
	return g
}
