package generator

import (
	"fmt"
	"testing"
)

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count += 2
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}

func TestLazyEvaluation(t *testing.T) {
	resume = integers()
	//fmt.Println(generateInteger())
	//fmt.Println(generateInteger())
	//fmt.Println(generateInteger())
	for {
		fmt.Println(generateInteger())

	}
}
