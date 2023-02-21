package dvargen

import "errors"

// ErrSubdir happens when directory contains sub directories.
var ErrSubdir = errors.New("directory should not contain sub directories")
