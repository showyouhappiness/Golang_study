package error_test

import (
	"errors"
	"testing"
)

//func GetFibonacci(n int) []int {
//	fibList := []int{1, 1}
//	for i := 2; i < n; i++ {
//		fibList = append(fibList, fibList[i-2]+fibList[i-1])
//	}
//	return fibList
//}
//
//func TestGetFibonacci(t *testing.T) {
//	t.Log(GetFibonacci(10))
//}

func GetFibonacci(n int) ([]int, error) {
	if n < 0 || n > 100 {
		return nil, errors.New("n should be in[2,100]")
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	//t.Log(GetFibonacci(10))
	if v, err := GetFibonacci(5); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}
