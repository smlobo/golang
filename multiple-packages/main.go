package main

import (
	"some-module/pkgA"
)

func main() {
	println("main()")
	println(pkgA.MainToAInterface(0))
}
