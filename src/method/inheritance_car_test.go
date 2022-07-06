/*
创建一个上面 Car 和 Engine 可运行的例子，并且给 Car 类型一个 wheelCount 字段和一个 numberOfWheels() 方法。
创建一个 Mercedes 类型，它内嵌 Car，并新建 Mercedes 的一个实例，然后调用它的方法。
然后仅在 Mercedes 类型上创建方法 sayHiToMerkel() 并调用它。
*/
package method

import (
	"fmt"
	"testing"
)

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
	Engine
}

// define a behavior for Car
func (car Car) numberOfWheels() int {
	return car.wheelCount
}

type Mercedes struct {
	Car //anonymous field Car
}

// a behavior only available for the Mercedes
func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi Angela!")
}

func (c *Car) Start() {
	fmt.Println("Car is started")
}

func (c *Car) Stop() {
	fmt.Println("Car is stopped")
}

func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	// drive to work
	c.Stop()
	// get out of car
}

func TestInheritanceCar(t *testing.T) {
	m := Mercedes{Car{4, nil}}
	fmt.Println("A Mercedes has this many wheels: ", m.numberOfWheels())
	m.GoToWorkIn()
	m.sayHiToMerkel()
}
