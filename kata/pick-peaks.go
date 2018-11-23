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
	p := 0

	for i := 1; i < len(array)-1; i++ {
		if array[i] > array[i-1] {
			p = i
		}
		if array[i] > array[i+1] && p != 0 {
			PosPeaks.Pos = append(PosPeaks.Pos, p)
			PosPeaks.Peaks = append(PosPeaks.Peaks, array[p])
			p = 0
		}
	}
	return PosPeaks
}
