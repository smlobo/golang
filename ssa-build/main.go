package main

import (
	"fmt"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"os"
)

func main() {
	// file as input
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <go-file>\n", os.Args[0])
		os.Exit(1)
	}

	config := packages.Config{Mode: packages.LoadAllSyntax}
	initial, err := packages.Load(&config, os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if packages.PrintErrors(initial) > 0 {
		fmt.Printf("Errors loading packages\n")
		os.Exit(1)
	}
	buildFlags := ssa.SanityCheckFunctions | ssa.InstantiateGenerics | ssa.NaiveForm
	//buildFlags := ssa.InstantiateGenerics | ssa.NaiveForm
	prog, _ := ssautil.AllPackages(initial, buildFlags)
	prog.Build()

	fmt.Printf("SSA build %s done!\n", os.Args[1])
}
