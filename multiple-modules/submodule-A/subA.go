package someLibraryA

import (
	"fmt"
	"runtime/debug"
)

func SomeSubA(x string) {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Printf("Failed to read build info")
	}
	fmt.Printf("SomeSubA(): [Got: %s] Module: %s\n", x, buildInfo.Path)
}

