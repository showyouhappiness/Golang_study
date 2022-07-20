package readWrite

import (
	"os"
	"testing"
)

func TestFileWrite(t *testing.T) {
	os.Stdout.WriteString("hello, world\n")
	f, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	f.WriteString("hello, world in a file\n")
}
