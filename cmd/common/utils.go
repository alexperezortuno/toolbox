package common

import (
	"fmt"
	"runtime"
	"strings"
)

func ParseDistro(input string, prefix string) string {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			return strings.TrimPrefix(line, prefix)
		}
	}
	return "Unknown"
}

func OsDetail() {
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Arch: %s\n", runtime.GOARCH)
	fmt.Printf("CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Go version: %s\n", runtime.Version())
}
