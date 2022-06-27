package array_test

import (
	"fmt"
	"testing"
)

func TestArrayLiterals(t *testing.T) {
	//var arrAge = [8]int{18, 20, 15, 22, 16}
	//var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrLazy = []int{5, 6, 7, 8, 22}	//注：初始化得到的实际上是切片slice
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}	//注：初始化得到的实际上是切片slice

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
}
func fp1(a *[3]int) { fmt.Println(a) }

func TestPointerArray(t *testing.T) {
	for i := 0; i < 3; i++ {
		fp1(&[3]int{i, i * i, i * i * i})
	}
}

const (
	WIDTH  = 1920
	HEIGHT = 1080
	// WIDTH =	5
	// HEIGHT = 4
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func TestMultidimArray(t *testing.T) {

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = pixel(x * y)
		}
	}
	fmt.Println(screen)

	//for row := range screen {
	//	for column := range screen[0] {
	//		screen[row][column] = 1
	//	}
	//}
	//
	//fmt.Println(screen)
}

func TestArraySum(t *testing.T) {
	array := [3]float64{7.0, 8.5, 9.1}
	print(&array)
	x := sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func sum(a *[3]float64) (sum float64) {
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
