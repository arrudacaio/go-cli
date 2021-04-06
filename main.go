package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	sliceArgs := os.Args[1:]
	data := sliceArgs[0]

	fmt.Println(args, sliceArgs, data)
}
