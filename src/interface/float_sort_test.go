package interface_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Float64Array []float64

func (p Float64Array) Len() int           { return len(p) }
func (p Float64Array) Less(i, j int) bool { return p[i] < p[j] }
func (p Float64Array) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func NewFloat64Array() Float64Array {
	return make([]float64, 25)
}

func (p Float64Array) Fill(n int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < n; i++ {
		p[i] = 100 * (rand.Float64())
	}
}

func (p Float64Array) List() string {
	s := "{ "
	for i := 0; i < p.Len(); i++ {
		if p[i] == 0 {
			continue
		}
		s += fmt.Sprintf("%3.1f ", p[i])
	}
	s += " }"
	return s
}

func (p Float64Array) String() string {
	return p.List()
}

func TestFloatSort(t *testing.T) {
	f1 := NewFloat64Array()
	f1.Fill(10)
	fmt.Printf("Before sorting %s\n", f1)
	Sort(f1)
	fmt.Printf("After sorting %s\n", f1)
	if IsSorted(f1) {
		fmt.Println("The float64 array is sorted!")
	} else {
		fmt.Println("The float64 array is NOT sorted!")
	}
}
