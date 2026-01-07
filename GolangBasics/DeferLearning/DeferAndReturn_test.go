package DeferLearning

import (
	"fmt"
	"testing"
)

func TestDeferAndReturn(t *testing.T) {
	DeferRun()
}

func TestDeferPrintArrTest(t *testing.T) {
	DeferPrintArrTest()
}

func TestOperateFinalResult(t *testing.T) {
	val := OperateFinalResult()
	fmt.Printf("OperateFinalResult: %v\n", val)
}
