package graph

import (
	"image"
	"image/color"
	"image/draw"
)

type Drawable interface {
	Draw(g Graphics)
}

// Graphics holds common methods for generating graphics.
// Based loosely on a similar construct in Java
type Graphics interface {
	Bounds() image.Rectangle
	Foreground(col color.Color) Graphics
	Background(col color.Color) Graphics

	Top() int
	Left() int
	Width() int
	Height() int

	Image() Image
	WritePNG(n string) error
	WriteJPG(n string) error

	Plot(x, y int) Graphics
	PlotPoint(p image.Point) Graphics

	Draw(p ...Drawable) Graphics

	DrawBetweenPoints(p1, p2 image.Point) Graphics
	DrawLine(x1, y1, x2, y2 int) Graphics

	DrawRect(x1, y1, x2, y2 int) Graphics
	DrawRectangle(rect image.Rectangle) Graphics
	FillRect(x1, y1, x2, y2 int) Graphics
	FillRectangle(rect image.Rectangle) Graphics

	// DrawImage is the "image/draw".Draw function
	DrawImage(r image.Rectangle, src image.Image, sp image.Point, op draw.Op) Graphics

	// DrawMask is the "image/draw".DrawMask function
	DrawMask(r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op draw.Op) Graphics

	// Filter applies a filter against the entire image being drawn
	Filter(filter Filter) error

	// Map applies a Mapper against the entire image being drawn
	Map(mapper Mapper) Graphics

	// Model applies a color.Model against the entire image being drawn
	Model(model color.Model) Graphics

	SetFont(font Font, size float64) Graphics
	TextSize(s string) Dimension
	DrawText(p image.Point, s string) Graphics
	DrawTextf(p image.Point, f string, a ...interface{}) Graphics
}
