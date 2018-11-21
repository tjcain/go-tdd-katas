package kata

// PosPeaks represents the position and value of peaks
type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// PickPeaks returns the positions and the values of the "peaks"
// (or local maxima) of a numeric array.
func PickPeaks(array []int) PosPeaks {
	PosPeaks := PosPeaks{[]int{}, []int{}}

	for i := len(array) - 1; i > 3; i-- {
		if array[i] != array[i-1] {
			break
		}
		array = array[:i]
	}

	for i := 1; i < len(array)-1; i++ {
		if array[i] > array[i-1] && array[i] >= array[i+1] {
			PosPeaks.Pos = append(PosPeaks.Pos, i)
			PosPeaks.Peaks = append(PosPeaks.Peaks, array[i])
		}
	}
	return PosPeaks
}
