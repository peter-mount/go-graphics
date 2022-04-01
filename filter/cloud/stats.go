package cloud

import (
	"math"
)

// Stats represents the calculated statistics from the filter
type Stats struct {
	total   int // Total pixels, should be >= visible as it can include no data entries
	visible int // Total pixels visible = cloud+sky
	cloud   int // Cloud pixels
	sky     int // Sky pixels
}

func percent(n, d int) float64 {
	if d == 0 {
		return 0
	}
	return 100.0 * float64(n) / float64(d)
}

// Cloud returns the percentage of the image representing clouds
func (s Stats) Cloud() float64 {
	return percent(s.cloud, s.visible)
}

// Sky returns the percentage of the image representing clear sky
func (s Stats) Sky() float64 {
	return percent(s.sky, s.visible)
}

// OKTA returns the OKTA level of the image
func (s Stats) OKTA() OKTA {
	// At least 1/8 of image contains data
	if float64(s.visible) > float64(s.total)/8 {
		return OKTA(math.Max(0, math.Min(8, math.Round(s.Cloud()/12.5))))
	}

	// Presume sky is obscured so OKTA 9
	return 9
}
