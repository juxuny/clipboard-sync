package lib

import (
	"testing"
)

func TestVector(t *testing.T) {
	v := Vector{0.6, 0.6, 0.6, 0.02, 0.01}
	num := 10000
	summaryMap := make(map[int]int)
	for i := 0; i < num; i++ {
		index, _ := v.DrawOne()
		summaryMap[index] += 1
	}
	t.Log(ToJSON(summaryMap))
}
