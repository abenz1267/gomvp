package internal

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Move(src, dest string) {
	module, root := getModulePath()

	if src == module {
		fixModule(src, dest, root)
		return
	}

	err := os.Chdir(root)
	if err != nil {
		panic(err)
	}

	base := strings.Split(dest, string(filepath.Separator))

	err = os.MkdirAll(filepath.Join(root, filepath.Join(base[:len(base)-1]...)), 0o755)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("git", "mv", src, dest)
	err = cmd.Run()
	if err != nil {
		err := os.RemoveAll(filepath.Join(root, dest))
		if err != nil {
			panic(err)
		}

		panic(err)
	}

	fixPackageName(src, dest, root)
	fixImports(src, dest, root, module)
}
