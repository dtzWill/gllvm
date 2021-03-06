package main

import (
	"github.com/SRI-CSL/gllvm/shared"
	"os"
)

func main() {

	shared.LogInfo("Entering %v\n", os.Args)
	// Parse command line
	args := os.Args
	args = args[1:]

	exitCode := shared.Compile(args, "clang")

	shared.LogInfo("Calling %v returned %v\n", os.Args, exitCode)

	//important to pretend to look like the actual wrapped command
	os.Exit(exitCode)

}
