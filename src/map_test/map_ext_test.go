package map_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true //添加元素，把元素设置为key，value设置为1
	n := 1          //1 is  existing
	//n := 3        //3 is not existing
	if mySet[n] { //判断元素是否存在
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}

	mySet[3] = true
	t.Log(len(mySet))

	delete(mySet, 1)
	t.Log(len(mySet))
	n = 1         //1 is not existing
	if mySet[n] { //判断元素是否存在
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
