package draw

import (
	"github.com/peter-mount/go-graphics"
	"image"
)

type Polygon interface {
	graph.Drawable
	Add(x, y int) Polygon
	AddPoint(pt image.Point) Polygon
	AddAll(pts ...image.Point) Polygon
	Bounds() image.Rectangle
}

func NewPolygon(closed bool) Polygon {
	return &polygon{closed: closed}
}

type polygon struct {
	points []image.Point
	bounds image.Rectangle
	closed bool
}

func (p *polygon) Add(x, y int) Polygon {
	return p.AddPoint(image.Point{X: x, Y: y})
}

func (p *polygon) AddPoint(pt image.Point) Polygon {
	p.points = append(p.points, pt)
	p.bounds.Add(pt)
	return p
}

func (p *polygon) AddAll(pts ...image.Point) Polygon {
	for _, pt := range pts {
		p.AddPoint(pt)
	}
	return p
}

func (p *polygon) Bounds() image.Rectangle {
	return p.bounds
}

func (p *polygon) Draw(g graph.Graphics) {
	var lp image.Point
	switch len(p.points) {
	// Empty Polygon
	case 0:
		return

		// Single point
	case 1:
		g.PlotPoint(p.points[0])

		// Draw the line
	default:
		for i, pt := range p.points {
			if i > 0 {
				g.DrawBetweenPoints(lp, pt)
			}
			lp = pt
		}

		// Close it
		if p.closed {
			g.DrawBetweenPoints(lp, p.points[0])
		}
	}
}
