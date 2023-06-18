package internal

import "errors"

// не найден ключ
var ErrKeyNotFound = errors.New("key not found")

const (
	keyNotFoundTemplate = "%w, key = %s"
)
