package xcode

import "strconv"

type XCode interface {
	Error() string
	Code() int
	Message() string
}

type Code int

func (c Code) Error() string {
	return strconv.Itoa(int(c))
}

func (c Code) Code() int {
	return int(c)
}

func (c Code) Message() string {
	return c.Error()
}

func Equal(code XCode, err error) bool {
	return false
}

func New(code int) Code {
	return Code(code)
}

func add(code int) Code {
	return Code(code)
}
