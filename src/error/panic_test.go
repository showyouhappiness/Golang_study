package error

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}
