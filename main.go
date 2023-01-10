package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func findos() {
	opsystem = runtime.GOOS
	fmt.Print(opsystem, "\n")
}

func command() {
	cmd := exec.Command("ls", "./")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	} else {
		fmt.Println("Output: ", string(out))
	}
}

func main() {
	findos()
	command()
}
