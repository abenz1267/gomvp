package main

import (
	"os"

	"github.com/abenz1267/gomvp/internal"
)

func main() {
	args := os.Args
	internal.Move(args[1], args[2])
}
