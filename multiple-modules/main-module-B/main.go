package main

import (
	"fmt"
	"runtime/debug"

	sA "github.com/sheldon/submodule-A"
	sB "github.com/sheldon/submodule-B"
)

func main() {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Failed to read build info")
	}
	fmt.Printf("main(): Module: %s\n", buildInfo.Path)
	fmt.Printf("  Dependencies:\n")
	for i, dep := range buildInfo.Deps {
        fmt.Printf("    [%d]: %s\n", i, dep.Path)
    }

    sA.SomeSubA("zoo")
    sB.SomeSubB("moo")
}
