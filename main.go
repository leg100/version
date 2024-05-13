package main

import (
	"fmt"
	"runtime/debug"
)

var (
	// Build-time parameters set via -ldflags
	Version = "unknown"
)

func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		// < go v1.18
		return
	}
	if info.Main.Version == "" || info.Main.Version == "(devel)" {
		// bin not built using `go install ...`
		return
	}
	// bin built using `go install ...`
	Version = info.Main.Version
}

func main() {
	fmt.Println("Version is", Version)
}
