package readWrite

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFileCopy(t *testing.T) {
	srcFile := "E:\\result_images2\\00.02.20P195922W34V1I1A0S000081064507411M0N02212c07R1E0Othund.jpg"
	dstFile := "E:\\result_images2\\03.02.20P195922W34V1I1A0S000081064507411M0N02212c07R1E0Othund.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}
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
