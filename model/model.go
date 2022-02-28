package model

import "errors"

var (
	ErrPersonNotNil = errors.New("Person can not be nil")
	ErrIDDoestExist = errors.New("ID does not exist")
)
