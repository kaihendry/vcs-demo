package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	build, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("No build info available")
		return
	}
	fmt.Printf("version: %s\n", build.Main.Version)
}
