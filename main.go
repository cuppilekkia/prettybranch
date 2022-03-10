package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	prefix       = "AC" // Your initials
	branchLength = 40   // Max branch lenght
	separator    = "-"  // Strings separator
)

// Pretty Branch
// Creates and checkes out a branch given a random input string
// with this format [PREFIX]-[given-strings-in-kebab-case]
// ie. "My awesome branch" => "AC-My-awesome-branch"
//
// TODO: add a config file and replace constants
func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}
	args = append([]string{prefix}, args...)
	branchname := strings.Join(args, separator)

	if len(branchname) > branchLength {
		branchname = branchname[:branchLength-1]
	}

	err := exec.Command("git", "checkout", "-b", branchname).Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Branch created and checked out!")
}
