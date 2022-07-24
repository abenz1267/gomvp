package internal

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

const GOMOD = "go.mod"

func getModulePath() (string, string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	sep := string(filepath.Separator)
	parts := strings.Split(wd, sep)

	current := 0
	path := filepath.Join(parts[:len(parts)-current]...)
	file, err := os.Open(sep + filepath.Join(path, GOMOD))

	for file == nil {
		current++
		path = filepath.Join(parts[:len(parts)-current]...)
		filepath := filepath.Join(path, GOMOD)
		file, err = os.Open(sep + filepath)
	}

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var module string
	for scanner.Scan() {
		module = scanner.Text()
		break
	}

	return strings.Split(module, " ")[1], sep + path
}
