package errors

import "errors"

var (
	IncorrectPositionItem = errors.New("removing item number more than len slice")
	NegativeIndex         = errors.New("negative index removing item")
)
