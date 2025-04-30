package main

/*
#include "mylib.h"
#cgo LDFLAGS: -L${SRCDIR} mylib.a
*/
import "C"
import "fmt"

type go123Result = C.Result
type goAllResult struct {
	cgoField C.Result
}

func main() {
	var result123 go123Result
	result123 = C.myFunction()
	result123.print()

	var resultAll goAllResult
	resultAll.cgoField = C.myFunction()
	resultAll.print()
}

func (r *go123Result) print() {
	fmt.Println("Receiver: Valid Go 1.23 / Invalid Go 1.24")
	fmt.Printf("value1: %d\n", r.value1)
	fmt.Printf("value2: %.2f\n", r.value2)
}

func (r *goAllResult) print() {
	fmt.Println("Receiver: (Fix) Valid Go 1.23 & 1.24")
	fmt.Printf("value1: %d\n", r.cgoField.value1)
	fmt.Printf("value2: %.2f\n", r.cgoField.value2)
}
