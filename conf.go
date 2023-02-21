package dvargen

import "io/fs"

// Conf configures the generation process.
type Conf struct {
	Path string      // Path of the generated file.
	Perm fs.FileMode // Perm of the generated file.
}
