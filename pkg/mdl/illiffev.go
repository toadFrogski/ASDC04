package mdl

type IlliffeVector struct {
	sizes []int
	data  interface{}
}

func NewIlliffeVector(sizes []int) *IlliffeVector {
	if len(sizes) == 0 {
		panic("Cannot create an empty Illiffe vector.")
	}
	return &IlliffeVector{
		sizes: sizes,
		data:  makeNestedSlice(sizes),
	}
}

func makeNestedSlice(sizes []int) interface{} {
	if len(sizes) == 1 {
		return make([]int, sizes[0])
	}
	subSlice := make([]interface{}, sizes[0])
	for i := 0; i < sizes[0]; i++ {
		subSlice[i] = makeNestedSlice(sizes[1:])
	}
	return subSlice
}

func (v *IlliffeVector) Get(indices []int) int {
	if len(indices) != len(v.sizes) {
		panic("Invalid number of indices.")
	}
	if indices[0] >= v.sizes[0] {
		panic("Index out of range.")
	}
	if len(v.sizes) == 1 {
		return v.data.([]int)[indices[0]]
	}
	return v.subVector(indices[0]).Get(indices[1:])
}

func (v *IlliffeVector) Set(indices []int, value int) {
	if len(indices) != len(v.sizes) {
		panic("Invalid number of indices.")
	}
	if indices[0] >= v.sizes[0] {
		panic("Index out of range.")
	}
	if len(v.sizes) == 1 {
		v.data.([]int)[indices[0]] = value
	} else {
		v.subVector(indices[0]).Set(indices[1:], value)
	}
}

func (v *IlliffeVector) subVector(index int) *IlliffeVector {
	return &IlliffeVector{
		sizes: v.sizes[1:],
		data:  v.data.([]interface{})[index],
	}
}
