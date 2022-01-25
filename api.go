package api

import (
	"fmt"
)

type Error struct {
	ErrCode int
	ErrMsg  string
	RawErr  error
}

type Calculator interface {
	Add(a, b int) int

	Div(a, b int) (float32, *Error)
}

func Add(a, b int) int {
	return a + b
}

func Div(a, b int) (float32, *Error) {
	if b == 0 {
		return 0, &Error{
			ErrCode: 1,
			ErrMsg:  "cannot div by 0",
			RawErr:  fmt.Errorf("cannot div by 0"),
		}
	}

	return float32(a) / float32(b), nil
}
