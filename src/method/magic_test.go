package method

import (
	"fmt"
	"testing"
)

type Base1 struct{}

func (Base1) Magic() {
	fmt.Println("base magic")
}

func (self Base1) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base1
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func TestMagic(t *testing.T) {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
