package cloud

// OKTA is a unit used in Aviation to represent the sky condition.
// It is an integer between 0 and 8.
//
// 0    SKC Sky Clear
// 1..2 FEW Few Clounds
// 3..4 SCT Scattered Clouds
// 5..7 BKN Broken Clouds
// 8    OVC Overcast
// 9    OBS Obscured - not used in aviation, but we use it if no value could be issued
type OKTA int

// Level returns OKTA as an integer.
func (o OKTA) Level() int {
	if o < 0 {
		return 0
	}
	if o > 9 {
		return 9
	}
	return int(o)
}

func (o OKTA) String() string {
	if o == 0 {
		return "SKC"
	}
	if o < 3 {
		return "FEW"
	}
	if o < 5 {
		return "SCT"
	}
	if o < 8 {
		return "BKN"
	}
	if o == 8 {
		return "OVC"
	}
	return "OBS"
}
