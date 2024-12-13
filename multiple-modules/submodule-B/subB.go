package someLibraryB

import (
	"fmt"
	"runtime/debug"
)

func SomeSubB(x string) {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Printf("Failed to read build info")
	}
	fmt.Printf("SomeSubB(): [Got: %s] Module: %s\n", x, buildInfo.Path)
}

