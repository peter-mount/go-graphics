// Package cloud implements a colour Mapper which can detect clouds within an image.
//
// It does this by taking the ratio of Red / Blue of a pixel.
// The higher this ration is, the more likely the pixel represents cloud.
// The lower the value it's more likely to be clear sky.
//
// Pixels which are out of range are mapped to Black.
package cloud

import (
	"image/color"
)

// Mapper will adjust an image so that it's pixels are based on if cloud or clear sky
// is detectable.
type Mapper struct {
	DarkLim  uint32      // Values below this are out of range
	LightLim uint32      // Values above this are out of range
	CloudLim float64     // lower limit of R/B for cloud
	SkyLim   float64     // lower limit of R/B for clear sky. Below this is no data
	Sky      color.Color // Colour for clear sky
	Cloud    color.Color // Color for clouds
	stats    Stats       // Calculated statistics
}

// New returns a new cloud Mapper with appropriate defaults set
func New() *Mapper {
	return &Mapper{
		DarkLim:  10,   // less than 10 get ignored
		LightLim: 250,  // greater than this is set to be the Sun
		CloudLim: 0.8,  // R/B > this gets marked as cloud
		SkyLim:   0.01, // R/B less than this indicates no data, if greater it's clear sky
		Sky:      color.RGBA{B: 255, A: 255},
		Cloud:    color.RGBA{R: 200, G: 200, B: 200, A: 255},
	}
}

// Mapper implements graph.Mapper which will map image colours to Sky/Cloud/Black
// whilst also updating this instance's statistics
func (c *Mapper) Mapper(col color.Color) color.Color {
	r, _, b, _ := col.RGBA()

	// Prevent divide-by-zero if no blue component
	if b > 0 {
		c.stats.total++

		rb := float64(r) / float64(b)

		if rb > c.CloudLim {
			c.stats.cloud++
			c.stats.visible++
			return c.Cloud
		}

		if rb > c.SkyLim {
			c.stats.sky++
			c.stats.visible++
			return c.Sky
		}

	}

	return color.Black
}

// Stats returns the calculated statistics from the underlying image
func (c *Mapper) Stats() Stats {
	return c.stats
}
