package main

import (
	"fmt"
	"strconv"
)

func hexToBin(hex string) string {
	ui, _ := strconv.ParseUint(hex, 16, 64)
	return fmt.Sprintf("%016b", ui)
}

func main() {
	bin := hexToBin("054C")
}
