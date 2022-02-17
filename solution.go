package square

import (
	"log"
	"math"
	"strconv"
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

func CalcSquare(sideLen float64, sidesNum string) float64 {
	sidesNums, err := strconv.Atoi(sidesNum)
	if err != nil {
		log.Fatal(err)
	}
	if sidesNums == 0 {
		sqr := sideLen * math.Pi
		return sqr
	} else if sidesNums == 3 {
		perimetr := 3 * sideLen
		sqr := math.Sqrt(perimetr * (perimetr - sideLen) * (perimetr - sideLen) * (perimetr - sideLen))
		return sqr
	} else {
		sqr := sideLen * sideLen
		return sqr
	}
}
