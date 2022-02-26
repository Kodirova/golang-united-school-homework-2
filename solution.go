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

type SideInit int

const (
	SidesTriangle SideInit = 3
	SidesSquare   SideInit = 4
	SidesCircle   SideInit = 0
)

func CalcSquare(sideLen float64, sidesNum SideInit) (area float64) {
	switch sidesNum {
	case SidesTriangle:
		area = (math.Pow(sideLen, 2.0) * math.Sqrt(3.0)) / 4.0
	case SidesSquare:
		area = math.Pow(sideLen, 2.0)
	case SidesCircle:
		area = math.Pi * math.Pow(sideLen, 2.0)
	default:
		area = 0
	}
	return area
}
