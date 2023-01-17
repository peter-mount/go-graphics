package graphics

import (
	"github.com/peter-mount/go-graphics"
	"github.com/peter-mount/go-graphics/util"
	"image"
	"image/color"
)

func New(img graph.Image) graph.Graphics {
	return &graphics{
		img:        img,
		bounds:     img.Bounds(),
		foreground: color.White,
		background: color.Transparent,
	}
}

func NewRect(rect image.Rectangle) graph.Graphics {
	return New(image.NewRGBA(rect))
}

type graphics struct {
	img        graph.Image
	bounds     image.Rectangle
	foreground color.Color
	background color.Color
	font       graph.Font
	fontSize   float64
}

func (g *graphics) Foreground(col color.Color) graph.Graphics {
	g.foreground = col
	return g
}

func (g *graphics) Background(col color.Color) graph.Graphics {
	g.background = col
	return g
}

func (g *graphics) Filter(filter graph.Filter) error {
	return filter.DoOver(g.img)
}

func (g *graphics) FilterBounds(filter graph.Filter, bounds image.Rectangle) error {
	return filter.Do(g.img, g.img, bounds)
}

func (g *graphics) Map(mapper graph.Mapper) graph.Graphics {
	mapper.DoOver(g.img)
	return g
}

func (g *graphics) MapBounds(mapper graph.Mapper, bounds image.Rectangle) graph.Graphics {
	mapper.Do(g.img, g.img, bounds)
	return g
}

func (g *graphics) Model(model color.Model) graph.Graphics {
	return g.Map(model.Convert)
}

func (g *graphics) WritePNG(n string) error {
	return util.WritePNG(n, g.img)
}

func (g *graphics) WriteJPG(n string) error {
	return util.WriteJPG(n, g.img)
}
