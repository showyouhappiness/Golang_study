package series

import "fmt"

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

/*想要从外部可以访问，那么声明的首字母必须是大写，如果是小写就不能在外部访问*/
func square(n int) int {
	return n * n
}

func GetFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
