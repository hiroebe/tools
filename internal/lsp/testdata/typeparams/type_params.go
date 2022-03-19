//go:build go1.18
// +build go1.18

package typeparams

func one[a int | string]()            {}
func two[a int | string, b float64 | int]() {}

func _() {
	one[]() //@rank("]", string, float64)
	two[]() //@rank("]", int, float64)
	two[int, f]() //@rank("]", float64, float32)
}

func slices[a []int | []float64]() {} //@item(tpInts, "[]int", "[]int", "type"),item(tpFloats, "[]float64", "[]float64", "type")

func _() {
	slices[]() //@rank("]", tpInts),rank("]", tpFloats)
}

type s[a int | string] struct{}

func _() {
	s[]{} //@rank("]", int, float64)
}
