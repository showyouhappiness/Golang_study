package readWrite

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestOsArgs(t *testing.T) {
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
