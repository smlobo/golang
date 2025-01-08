package pkgB

import (
	"some-module/pkgC"
)

func AToBInterface(x int) int {
	println("AToBInterface()")
	x += 2
	return pkgC.BToCInterface(x)
}
