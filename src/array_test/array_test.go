package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
	arr1[1] = 5
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for index, value := range arr3 { //类似forEach
		t.Log(index, value)
	}
	for _, value := range arr3 { //使用_代表并不关心这个值的结果
		t.Log(value)
	}
}

func TestArraySection(t *testing.T) {
	arr4 := [...]int{1, 2, 3, 4, 5}
	arr4Sec := arr4[:] //Go的切片不支持负数位数，和Python不一样
	t.Log(arr4Sec)
}
