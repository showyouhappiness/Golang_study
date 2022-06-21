package client

import (
	"Golang_study/src/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
	t.Log(series.Square(5))
}
