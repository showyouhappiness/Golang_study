/*
一个将英文饮料名映射为中文的集合；先打印所有的饮料，然后打印原名和翻译后的名字。接下来按照英文名排序后再打印出来。
*/
package map_test

import (
	"fmt"
	"sort"
	"testing"
)

func TestDrinks(t *testing.T) {
	drinks := map[string]string{
		"beer":   "啤酒",
		"wine":   "葡萄酒",
		"water":  "水",
		"coffee": "咖啡",
		"thea":   "茶"}
	sdrinks := make([]string, len(drinks))
	ix := 0

	fmt.Printf("The following drinks are available:\n")
	for eng := range drinks {
		sdrinks[ix] = eng
		ix++
		fmt.Println(eng)
	}

	fmt.Println("")
	for eng, fr := range drinks {
		fmt.Printf("The chinese for %s is %s\n", eng, fr)
	}

	// SORTING:
	fmt.Println("")
	fmt.Println("Now the sorted output:")
	sort.Strings(sdrinks)

	fmt.Printf("The following sorted drinks are available:\n")
	for _, eng := range sdrinks {
		fmt.Println(eng)
	}

	fmt.Println("")
	for _, eng := range sdrinks {
		fmt.Printf("The chinese for %s is %s\n", eng, drinks[eng])
	}
}
