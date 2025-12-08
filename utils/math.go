package utils

import "math"

func Pow(n, p int) int {
	if p == 0 {
		return 1
	}
	base := n
	for i := 0; i < p-1; i++ {
		n *= base
	}
	return n
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func SquareEuclideanDistance(x1, x2, y1, y2, z1, z2 int) int {
	return Pow(x1-x2, 2) + Pow(y1-y2, 2) + Pow(z1-z2, 2)
}

func EuclideanDistance(x1, x2, y1, y2, z1, z2 int) float64 {
	return math.Sqrt(float64(SquareEuclideanDistance(x1, x2, y1, y2, z1, z2)))
}

func ManhattanDistance(x1, x2, y1, y2, z1, z2 int) int {
	return Abs(x1-x2) + Abs(y1-y2) + Abs(z1-z2)
}
