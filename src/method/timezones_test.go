/*
为 int 定义别名类型 TZ，定义一些常量表示时区，比如 UTC，定义一个 map，它将时区的缩写映射为它的全称，
比如：UTC -> "Universal Greenwich time"。为类型 TZ 定义 String() 方法，它输出时区的全称。
*/
package method

import (
	"fmt"
	"testing"
)

type TZ int

const (
	HOUR TZ = 60 * 60
	UTC  TZ = 0 * HOUR
	EST  TZ = -5 * HOUR
	CST  TZ = -6 * HOUR
)

var timeZones = map[TZ]string{
	UTC: "Universal Greenwich time",
	EST: "Eastern Standard time",
	CST: "Central Standard time"}

func (tz TZ) String() string { // Method on TZ (not ptr)
	if zone, ok := timeZones[tz]; ok {
		return zone
	}
	return ""
}

func TestTimezones(t *testing.T) {
	fmt.Println(EST) // Print* knows about method String() of type TZ
	fmt.Println(0 * HOUR)
	fmt.Println(-6 * HOUR)
}
