package extension2_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("", host)
}

type Dog struct {
	Pet
}

//func (d *Dog) Speak() {
//	fmt.Print("Wang Wang")
//}
//
//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Println("", host)
//}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("hi")
}
