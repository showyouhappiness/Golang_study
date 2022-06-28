package forrange

import (
	"fmt"
	"testing"
)

func TestSlicesForRange(t *testing.T) {
	slice1 := make([]int, 4)

	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
}
func TestSlicesForRange2(t *testing.T) {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
		println(item)
	}
}
