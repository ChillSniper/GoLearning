package DeferLearning

import (
	"fmt"
	"testing"
)

func TestCopyFile(t *testing.T) {
	wr, err := CopyFile("/AfterCopy.go", "/deferExecutionOrder.go")
	if err != nil {
		return
	}
	fmt.Println(wr)
}

func TestDealWithRecoverToPanic(t *testing.T) {
	DealWithRecoverToPanic()
}
