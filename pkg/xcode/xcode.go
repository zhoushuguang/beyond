package xcode

import "strconv"

type XCode interface {
	Error() string
	Code() int
	Message() string
	Details() []interface{}
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

func (c Code) Details() []interface{} {
	return nil
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
