package readWrite

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func TestReadInput(t *testing.T) {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go
}

var inputReader *bufio.Reader
var input1 string
var err error

func TestReadInput2(t *testing.T) {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input1, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input1)
	}
}
