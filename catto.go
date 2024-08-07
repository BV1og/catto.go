package main

type fundomentalTypes interface {
	int | float64 | string | bool | byte | int8 | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | uintptr | float32
}
type Numbers interface {
	int | float64 | byte | int8 | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | uintptr | float32
}
type fundomentalTypesWithArrays interface {
	int | float64 | string | bool | byte | int8 | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | uintptr | float32 |
		[]int | []float64 | []string | []bool | []byte | []int8 | []int16 | []int32 | []int64 | []uint | []uint16 | []uint32 | []uint64 | []uintptr | []float32 |
		[][]int | [][]float64 | [][]string | [][]bool | [][]byte | [][]int8 | [][]int16 | [][]int32 | [][]int64 | [][]uint | [][]uint16 | [][]uint32 | [][]uint64 | [][]uintptr | [][]float32 |
		[][][]int | [][][]float64 | [][][]string | [][][]bool | [][][]byte | [][][]int8 | [][][]int16 | [][][]int32 | [][][]int64 | [][][]uint | [][][]uint16 | [][][]uint32 | [][][]uint64 | [][][]uintptr | [][][]float32
}

func Contains[T fundomentalTypes](m []T, s T) bool {
	for _, v := range m {
		if s == v {
			return true
		}
	}
	return false
}

func Remove[T fundomentalTypes](m []T, i int) []T {
	temp := m[0 : i-1]
	for _, v := range m[i:len(m)] {
		temp = append(temp, v)
	}
	return temp
}

func Maximum[T Numbers](m []T) T {
	max := T(0)
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func Minimum[T Numbers](m []T) T {
	min := m[0]
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	return min
}

func GenerateArray[T fundomentalTypesWithArrays](e T, s int) []T {
	r := []T{}
	for i := 0; i < s; i++ {
		r = append(r, e)
	}
	return r
}

/*func main() {
	a := GenerateArray(GenerateArray(GenerateArray(GenerateArray(0, 5), 5), 5), 5)
	fmt.Println(a)
}*/

func main() {}
