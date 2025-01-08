package pkgA

import (
	"some-module/pkgB"
)

func MainToAInterface(x int) int {
	println("MainToAInterface()")
	x += 1
	return pkgB.AToBInterface(x)
}
