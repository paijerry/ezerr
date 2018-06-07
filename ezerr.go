package ezerr

import "fmt"

// EzErr -
type EzErr struct {
	code   string
	msg    string
	detail []interface{}
}

// Error - as an error, must have this method
func (e *EzErr) Error() string {
	return fmt.Sprintf("%v: %v", e.code, e.msg)
}

// GetCode - get code of the ezerr
func (e *EzErr) GetCode() string {
	return e.code
}

// GetMsg - get message of the ezerr
func (e *EzErr) GetMsg() string {
	return e.msg
}

// SetDetail - set any detail into ezerr
func (e *EzErr) SetDetail(detail string) {
	e.detail = append(e.detail, detail)
}

// GetDetail - get all the detail of the ezerr
func (e *EzErr) GetDetail() []interface{} {
	return e.detail
}

// GenEzErr - generate ezerr
func GenEzErr(code string, msg string) *EzErr {
	return &EzErr{
		code: code,
		msg:  msg,
	}
}

// IsEz - check the error is ezerr or not
func IsEz(err error) bool {
	_, ok := err.(*EzErr)
	return ok
}
