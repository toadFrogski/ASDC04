package mdl

type Interval struct {
	start int
	end   int
}

func NewInterval(interval []int) *Interval {
	return &Interval{start: interval[0], end: interval[1]}
}

type DefiningVector struct {
	intervals         []Interval
	dimension         int
	dimensionSizes    []int
	size              int
	dimensionSizesSum int
}

func NewDefiningVector(dimensions []int, intervals [][]int) *DefiningVector {
	d := &DefiningVector{nil, 0, nil, 0, 0}
	for _, interval := range intervals {
		d.addInterval(interval)
	}
	size := 1
	for i, dim := range dimensions {
		d.dimensionSizes = append(d.dimensionSizes, size)
		d.dimensionSizesSum += size * d.intervals[i].start
		size *= dim
	}
	d.size = size
	d.dimension = len(d.dimensionSizes)
	return d
}

func (d *DefiningVector) addInterval(interval []int) {
	d.intervals = append(d.intervals, *NewInterval(interval))
}
