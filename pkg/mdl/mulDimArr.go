package mdl

import (
	"math/rand"
	"time"
)

type MultidimensionalArray struct {
	dimensions []int
	data       []int
	display    *IlliffeVector
	defv       *DefiningVector
}

func NewMultidimensionalArray(dimensions []int, intervals [][]int) *MultidimensionalArray {
	size := 1
	for _, dim := range dimensions {
		size *= dim
	}

	display := NewIlliffeVector(dimensions)
	defv := NewDefiningVector(dimensions, intervals)

	data := make([]int, size)
	m := &MultidimensionalArray{
		dimensions: dimensions,
		data:       data,
		display:    display,
		defv:       defv,
	}

	rand.Seed(time.Now().UnixNano())
	for i := range m.data {
		m.data[i] = rand.Int()
	}

	for i := range m.data {
		m.display.Set(getCoords(i, dimensions), m.data[i])
	}

	return m
}

func (m *MultidimensionalArray) GetDirect(indexes []int) int {
	defer timer("getDirect")()
	offset, grade, len := 0, 1, len(indexes)
	for i := 0; i < len; i++ {
		offset += indexes[len-i-1] * grade
		grade *= m.dimensions[len-i-1]
	}
	return m.data[offset]
}

func (m *MultidimensionalArray) GetIlliffe(indexes []int) int {
	defer timer("getIlliffe")()
	return m.display.Get(indexes)
}

func (m *MultidimensionalArray) GetDefining(indexes []int) int {
	defer timer("getDefining")()
	offset, len := 0, len(indexes)
	for i := 0; i < len; i++ {
		offset += indexes[i] * m.defv.dimensionSizes[m.defv.dimension-i-1]
	}
	offset -= m.defv.dimensionSizesSum
	return m.data[offset]
}
