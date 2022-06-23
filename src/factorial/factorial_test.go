// 阶乘
package factorial

import (
	"fmt"
	"math/big"
	"testing"
)

// 用普通的int64来计算阶乘，最多只能计算到20的阶乘，就溢出了
func TestFactorial(t *testing.T) { // 递归 阶乘
	i := 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func TestInt(t *testing.T) {
	var i int64
	for i = 1; i <= 40; i++ {
		fmt.Println(aInt(i))
	}
}

func aInt(s int64) int64 {
	if s == 1 {
		return 1
	} else {
		return s * aInt(s-1)
	}
}

func TestBigNum(t *testing.T) {
	var i int64
	for i = 1; i <= 40; i++ {
		fmt.Println(aBig(big.NewInt(i)))
	}
}

func aBig(s *big.Int) *big.Int {
	if s.Int64() == 1 {
		return big.NewInt(1)
	} else {
		return s.Mul(s, aBig(big.NewInt(s.Int64()-1)))
	}
}
