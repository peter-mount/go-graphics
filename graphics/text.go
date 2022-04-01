package graphics

import (
	"fmt"
	"github.com/peter-mount/go-graphics"
	"github.com/peter-mount/go-graphics/draw"
	"image"
)

func (g *graphics) SetFont(font graph.Font, size float64) graph.Graphics {
	g.font = font
	g.fontSize = size
	return g
}

func (g *graphics) textSize(s string) (float64, float64) {
	w, h := 0, 0
	for _, c := range s {
		w1, h1 := g.font.Size(c)
		w = w + w1
		if h1 > h {
			h = h1
		}
	}
	return float64(w), float64(h)
}

func (g *graphics) TextSize(s string) graph.Dimension {
	w, h := g.textSize(s)
	return graph.NewDimension(int(w*g.fontSize), int(h*g.fontSize))
}

func (g *graphics) DrawText(p image.Point, s string) graph.Graphics {
	tw, th := g.textSize(s)
	tb := image.Rect(0, 0, int(tw), int(th))
	ti := image.NewRGBA(tb)

	pt := image.Point{}
	for _, c := range s {
		fmt.Printf("%3d %c %v\n", c, c, pt)
		pt = g.font.DrawRune(ti, pt, c)
	}

	fs := g.fontSize / th
	dp := int(fs)
	var r image.Rectangle
	for y := tb.Min.Y; y < tb.Max.Y; y++ {
		for x := tb.Min.X; x < tb.Max.X; x++ {
			_, _, _, a := ti.At(x, y).RGBA()
			if a > 0 {
				r.Min.X = int(float64(p.X) + (fs * float64(x)))
				r.Min.Y = int(float64(p.Y) + (fs * float64(y)))
				r.Max.X = r.Min.X + dp
				r.Max.Y = r.Min.Y + dp
				draw.FillRect(g.img, r, g.foreground)
			}
		}
	}

	//draw.Draw(g.img, ts.Rect(p.X, p.Y), ti, image.Point{}, draw.Over)

	return g
}

func (g *graphics) DrawTextf(p image.Point, f string, a ...interface{}) graph.Graphics {
	return g.DrawText(p, fmt.Sprintf(f, a...))
}
