package repository

import "errors"

var ErrBookNotFound = errors.New("book not found")
var ErrAuthorNotFound = errors.New("author not found")