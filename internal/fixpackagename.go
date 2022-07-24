package internal

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
)

func fixPackageName(src, dest, root string) {
	oldpkgparts := strings.Split(src, "/")
	oldpkg := oldpkgparts[len(oldpkgparts)-1]
	newpkgparts := strings.Split(dest, "/")
	newpkg := newpkgparts[len(newpkgparts)-1]
	newdest := filepath.Join(root, dest)

	info, err := os.ReadDir(newdest)
	if err != nil {
		panic(err)
	}

	for _, v := range info {
		if v.IsDir() {
			continue
		}

		path := filepath.Join(newdest, v.Name())
		file, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		replaced := bytes.ReplaceAll(file, []byte("package "+oldpkg), []byte("package "+newpkg))
		err = os.WriteFile(path, replaced, v.Type().Perm())
		if err != nil {
			panic(err)
		}
	}
}
