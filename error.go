package torsaver

import "fmt"

type errWrap struct {
	err error
	msg string
}

func (e errWrap) Error() string {
	return fmt.Sprintf("%v:%v", e.msg, e.err)
}

func Wrap(e error, msg string) error {
	return &errWrap{
		err: e,
		msg: msg,
	}
}
