package dvargen

import (
	"os"
	"testing"
)

func TestDVarGen(t *testing.T) {
	dname, err := os.MkdirTemp("", "dvargen")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dname)
	// dname = "testdata"

	t.Run("Generate", func(t *testing.T) {
		// Will create file in the current directory.

		// vDesc := DVarDesc{
		// 	Dir:     "testdata/dir1",
		// 	Varname: "m",
		// 	Package: "dvargen",
		// }
		// err := Generate(vDesc)
		// if err != nil {
		// 	t.Fatal(err)
		// }
	})

	t.Run("GenerateAs", func(t *testing.T) {
		vDesc := DVarDesc{
			Dir:     "testdata/dir1",
			Varname: "m",
			Package: "testdata",
		}
		conf := Conf{
			Path: dname,
			Perm: os.ModePerm,
		}
		err := GenerateAs(vDesc, conf)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := os.Stat(filepath(conf.Path, vDesc)); err != nil {
			t.Error("file was not generated")
		}
	})

	t.Run("With subdir", func(t *testing.T) {
		vDesc := DVarDesc{
			Dir:     "testdata/dir2",
			Varname: "m",
			Package: "testdata",
		}
		err := Generate(vDesc)
		if err != ErrSubdir {
			t.Fatal(err)
		}
	})

}
