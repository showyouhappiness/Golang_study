package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4, 5}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2]) //因为上面定义的是三位长度，这里只能打印三位，超过会报错。
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	month := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q1 := month[3:6]
	t.Log(Q1, len(Q1), cap(Q1)) //len是取得个数，cap是从第一个开始取然后到结尾
	Q2 := month[5:8]
	t.Log(Q2, len(Q2), cap(Q2)) //len是取得个数，cap是从第一个开始取然后到结尾
	Q2[0] = "Unknown"
	t.Log(Q1) //在这个共享的空间里，Q1和Q2都是一个数据，修改一个那么另外一个就会发生改变，这个如果使用append就不会这样
	t.Log(month)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{1, 2, 3, 4, 5, 6}
	//if a == b { // slice只能和nil进行比较其他的都会报错
	//	t.Log("equal")
	//}
	t.Log(a, b)
}
