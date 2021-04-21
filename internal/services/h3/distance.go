package h3

import (
	"math"

	"github.com/uber/h3-go"
)

func DistanceBetweenH3Indices(a int64, b int64) float64 {
	const earthRadius = 6378136

	hSin := func(theta float64) float64 {
		return math.Pow(math.Sin(theta/2), 2)
	}

	pointA := h3.ToGeo(h3.H3Index(a))
	pointB := h3.ToGeo(h3.H3Index(b))

	var aLat, aLon, bLat, bLon float64
	aLat = pointA.Latitude * math.Pi / 180
	aLon = pointA.Longitude * math.Pi / 180
	bLat = pointB.Latitude * math.Pi / 180
	bLon = pointB.Longitude * math.Pi / 180

	h := hSin(bLat-aLat) + math.Cos(aLat)*math.Cos(bLat)*hSin(bLon-aLon)
	return 2 * earthRadius * math.Asin(math.Sqrt(h))
}
