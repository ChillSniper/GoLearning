package ErrorLearning

import (
	"errors"
	"fmt"
	"testing"
)

func TestGetPositiveSelfAdd(t *testing.T) {
	num1, err1 := GetPositiveSelfAdd(1)
	fmt.Printf("num is %d, err is %v\n", num1, err1)

	num2, err2 := GetPositiveSelfAdd(-2)
	fmt.Printf("num is %d, err is %v\n", num2, err2)

	err3 := errors.New("hello")
	err4 := errors.New("hello")
	fmt.Println(errors.Is(err4, err3))

	fmt.Println(err3.Error())
	fmt.Println(err3.Error() == err4.Error())

}
