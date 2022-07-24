package internal

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func fixImports(oldpkg, newpkg, root, module string) {
	oldusageparts := strings.Split(oldpkg, "/")
	oldusage := oldusageparts[len(oldusageparts)-1] + "."
	newusageparts := strings.Split(newpkg, "/")
	newusage := newusageparts[len(newusageparts)-1] + "."

	oldpkg = module + "/" + oldpkg
	newpkg = module + "/" + newpkg

	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(path, ".go") {
				file, err := os.ReadFile(path)
				if err != nil {
					panic(err)
				}

				replaced := bytes.ReplaceAll(file, []byte(oldpkg), []byte(newpkg))
				replaced = bytes.ReplaceAll(replaced, []byte(oldusage), []byte(newusage))

				err = os.WriteFile(path, replaced, info.Mode().Perm())
				if err != nil {
					panic(err)
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
