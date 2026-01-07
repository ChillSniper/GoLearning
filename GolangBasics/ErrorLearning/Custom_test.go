package ErrorLearning

import (
	"fmt"
	"testing"
)

func TestCustom(t *testing.T) {
	err := NewError(100, "test MyError")
	fmt.Printf("code is %d, msg is %s", Code(err), Msg(err))
}
