package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

var SidesTriangle = 3
var SidesSquare = 4
var SidesCircle = 0

func CalcSquare(sideLen float64, sides int) float64 {

	if sides == 0 {
		sqr := sideLen * math.Pi
		return sqr
	} else if sides == 3 {
		sqr := 0.5 * sideLen * sideLen * (math.Sqrt(3) / 2)
		return sqr
	} else {
		sqr := sideLen * sideLen
		return sqr
	}
}
