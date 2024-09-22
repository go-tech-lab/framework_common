package util

import (
	"fmt"
	"runtime"
)

func PrintPanicStack() {
	var buf [9192]byte
	n := runtime.Stack(buf[:], false)
	fmt.Errorf("==> %s\n", string(buf[:n]))
}
