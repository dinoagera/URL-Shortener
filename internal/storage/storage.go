package storage

import "errors"

var (
	ErrNotFound  = errors.New("URL not found")
	ErrUrlExists = errors.New("URL exists")
)
