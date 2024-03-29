package error_test

import (
	"errors"
	"fmt"
	"testing"
)

func GetFibonacci(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}

func TestGetFibonacci(t *testing.T) {
	t.Log(GetFibonacci(10))
}

func GetFibonacci1(n int) ([]int, error) {
	if n < 0 || n > 100 { //错误处理
		return nil, errors.New("n should be in[2,100]") //返回错误
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci1(t *testing.T) {
	//t.Log(GetFibonacci(10))
	if v, err := GetFibonacci1(10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

var errNotFound = errors.New("not found error")

func TestError(t *testing.T) {
	fmt.Printf("error: %v\n", errNotFound)
	if _, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s", err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}
	return f, nil
	// implementation of Sqrt
}
