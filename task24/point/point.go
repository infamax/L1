package point

import "math"

type Point struct {
	x float64
	y float64
}

func New(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(point1, point2 *Point) float64 {
	return math.Sqrt(math.Pow(point1.x-point2.x, 2) +
		math.Pow(point1.y-point2.y, 2))
}
