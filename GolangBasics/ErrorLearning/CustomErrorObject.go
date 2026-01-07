package ErrorLearning

import "fmt"

type MyError struct {
	code int
	msg  string
}

func (m MyError) Error() string {
	return fmt.Sprintf("code:%d,msg:%v", m.code, m.msg)
}

func NewError(code int, msg string) error {
	return MyError{
		code: code,
		msg:  msg,
	}
}

func Code(err error) int {
	if e, ok := err.(MyError); ok {
		return e.code
	}
	return -1
}

func Msg(err error) string {
	if e, ok := err.(MyError); ok {
		return e.msg
	}
	return ""
}
