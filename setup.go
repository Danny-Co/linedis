package main

import (
	"fmt"
	"runtime"
)

var opsystem string

func findos() {
	opsystem = runtime.GOOS
	fmt.Print(opsystem)
}
