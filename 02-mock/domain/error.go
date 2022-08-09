package domain

import "errors"

var (
	ErrRecordNotFound   = errors.New("record not found")
	ErrConstraintUnique = errors.New("constraint unique")
)
