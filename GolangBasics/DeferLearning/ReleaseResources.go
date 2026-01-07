package DeferLearning

import (
	"fmt"
	//"fmt"
	"io"
	"os"
)

func CopyFile(dstFile, srcFile string) (wr int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		return
	}
	defer dst.Close()
	wr, err = io.Copy(dst, src)
	return wr, err
}

func DealWithRecoverToPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	a := 1
	b := 0
	fmt.Printf("result: %v\n", a/b)
}
