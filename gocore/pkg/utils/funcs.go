package utils

import (
	"math"
	"ws-core/internal/common"
)

const PI = 3.1415926535898

func Dist(p1, p2 common.Coord) int {
	if p1.Lat == p2.Lat && p1.Long == p2.Long {
		return 0
	}
	R := 6371.0

	x1 := p1.Lat - p2.Lat
	dLat := x1 * PI / 180.0
	x2 := p1.Long - p2.Long
	dLon := x2 * PI / 180.0
	A := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(p1.Lat*PI/180.0)*math.Cos(p2.Lat*PI/180.0)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(A), math.Sqrt(1-A))
	d := R * c
	return int(d * 1.5 * 1000)
}
