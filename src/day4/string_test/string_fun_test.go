package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)
	for _, part := range parts {
		t.Logf(part)
	}
	t.Log(strings.Join(parts, ""))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	//t.Log(10 + strconv.Atoi("10"))
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
