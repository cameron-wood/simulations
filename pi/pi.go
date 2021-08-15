package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("π ≈", calculatePi(100_000_000))
}

func calculatePi(n int) float64 {
	const radius = 1
	var innerPoints int

	for i := 0; i < n; i++ {
		x, y := rand.Float64(), rand.Float64()
		if math.Sqrt(x*x+y*y) < radius {
			innerPoints++
		}
	}

	return 4 * float64(innerPoints) / float64(n)
}
