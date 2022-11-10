package map_test

import (
	"fmt"
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1 = %d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2 = %d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3 = %d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	/*
		两个输出的都是0，
		一个不存在的key输出的是0，
		而另一个存在的key本身的值是0，输出也就是0
		无法区分到底是不存在还是它本身的值就是0
	*/
	//下面的代码就可以区分，分别是给值和不给值
	//m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		fmt.Println(k, v)
	}

}

func TestIsPresent(t *testing.T) {
	map1 := map[string]string{
		"a": "apple",
		"b": "banana",
	}
	v, isPresent := map1["b"]

	if isPresent {
		fmt.Print("The key is present in map.", v)
	} else {
		fmt.Print("The key is not present in map.", v)
	}
}
