package dvargen

// DVarDesc is te description of the generated var.
type DVarDesc struct {
	Dir     string // Generates var from content of this directory.
	Varname string // Name of the generated var, also name of the generated
	// file.
	Package string // Package of the generated file.
}
