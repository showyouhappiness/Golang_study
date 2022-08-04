package readWrite

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

var numberFlag = flag.Bool("n", false, "number each line")

func cat1(r *bufio.Reader) {
	i := 1
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func TestCat(t *testing.T) {
	flag.Parse()
	fmt.Println("flag.Args(): ", flag.Args())
	fmt.Println("flag.NArg(): ", flag.NArg())
	if flag.NArg() == 0 {
		cat1(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat1(bufio.NewReader(f))
	}
}
