package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/abenz1267/gomvp/internal"
)

func main() {
	checkGitConditions()

	args := os.Args
	internal.Move(args[1], args[2])
}

func checkGitConditions() {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")

	err := cmd.Run()
	if err != nil {
		log.Println(err)
		log.Fatal("Are you running gomvp inside a git directory?")
	}

	cmd = exec.Command("git", "status", "--porcelain")

	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	if string(out) != "" {
		log.Fatal("Make sure there are no uncommited changes before running gomvp!")
	}
}
