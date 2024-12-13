package main

/*
#include "mylib.h"
#cgo LDFLAGS: -L${SRCDIR} mylib.a
*/
import "C"
import "fmt"

func main() {
    result := C.myFunction()
    fmt.Printf("value1: %d\n", result.value1)
    fmt.Printf("value1: %.2f\n", result.value2)
    fmt.Printf("value1: %v\n", result.v3)
    fmt.Printf("value1: %v\n", result.v4)
}