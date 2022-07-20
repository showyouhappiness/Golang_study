package readWrite

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFileCopy(t *testing.T) {
	file, err := CopyFile("target.txt", "source.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
