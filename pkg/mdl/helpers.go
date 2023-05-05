package mdl

import (
	"fmt"
	"time"
)

func getCoords(index int, shape []int) []int {
	coords := make([]int, len(shape))
	for i := len(shape) - 1; i >= 0; i-- {
		coords[i] = index % shape[i]
		index = index / shape[i]
	}
	return coords
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
