package main

import (
	"fmt"
	"runtime"
)

func main() {
	// print OS and ARCH
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
