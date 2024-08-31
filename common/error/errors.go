package error

import "errors"

// Define common errors
var (
    ErrNotFound = errors.New("resource not found")
    ErrInternal = errors.New("internal server error")
)
