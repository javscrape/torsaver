package torsaver

import "fmt"

type errWrap struct {
	err error
	msg string
}

// Error ...
func (e errWrap) Error() string {
	return fmt.Sprintf("%v:%v", e.msg, e.err)
}

// Wrap ...
func Wrap(e error, msg string) error {
	return &errWrap{
		err: e,
		msg: msg,
	}
}
