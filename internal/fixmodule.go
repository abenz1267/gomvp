package internal

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func fixModule(src, dest, root string) {
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "go.mod") {
				file, err := os.ReadFile(path)
				if err != nil {
					panic(err)
				}

				replaced := bytes.ReplaceAll(file, []byte(src), []byte(dest))

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
