package text

import (
	"github.com/peter-mount/go-graphics"
	"image"
	"image/color"
)

type Text struct {
	font       graph.Font
	size       int
	pos        image.Point
	foreground color.Color
	background color.Color
}

func New(font graph.Font, size int) *Text {
	return &Text{
		font:       font,
		size:       size,
		foreground: color.White,
		background: color.Transparent,
	}
}

func (t *Text) SetSize(size int) *Text {
	t.size = size
	return t
}

func (t *Text) SetPoint(p image.Point) *Text {
	t.pos = p
	return t
}

func (t *Text) SetPos(x, y int) *Text {
	return t.SetPoint(image.Point{X: x, Y: y})
}

func (t *Text) Foreground(c color.Color) *Text {
	t.foreground = c
	return t
}

func (t *Text) Background(c color.Color) *Text {
	t.foreground = c
	return t
}

func (t *Text) TextSize(s string) (int, int) {
	w, h := 0, 0
	for _, c := range s {
		w1, h1 := t.font.Size(c)
		w = w + w1
		if h1 > h {
			h = h1
		}
	}
	return w, h
}

func (t *Text) Draw(img graph.Image, s string) *Text {
	w, h := t.TextSize(s)
	b := image.Rect(0, 0, w, h)
	ni := image.NewRGBA(b)

	p := t.pos
	for _, c := range s {
		p = t.font.DrawRune(ni, p, c)
	}

	return t.SetPoint(p)
}
