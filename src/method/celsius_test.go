/*
为 float64 定义一个别名类型 Celsius，并给它定义 String()，它输出一个十进制数和 °C 表示的温度值。
*/
package method

import (
	"fmt"
	"strconv"
	"testing"
)

type Celsius float64

func (c Celsius) String() string {
	return "The temperature is: " + strconv.FormatFloat(float64(c), 'f', 1, 32) + " °C"
}

func TestCelsius(t *testing.T) {
	var c Celsius = 18.36
	fmt.Println(c)
}
