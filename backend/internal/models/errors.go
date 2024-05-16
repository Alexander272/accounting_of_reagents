package models

import "errors"

var (
	ErrNoRows   = errors.New("row not found")
	ErrBadValue = errors.New("bad value")

	ErrSessionEmpty = errors.New("user session not found")
)
