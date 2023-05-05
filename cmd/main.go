package main

import (
	"fmt"

	. "github.com/toadFrogski/ASDC04/pkg/mdl"
)

func main() {
	shape := []int{3, 3}
	intervals := [][]int{{0, 2}, {0, 2}}
	multiDimArray := NewMultidimensionalArray(shape, intervals)
	res1 := multiDimArray.GetDirect([]int{1, 2})
	fmt.Printf("%v\n", res1)
	res2 := multiDimArray.GetIlliffe([]int{1, 2})
	fmt.Printf("%v\n", res2)
	res3 := multiDimArray.GetDefining([]int{1, 2})
	fmt.Printf("%v\n", res3)
}
