package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	isSuccess := args[0]
	if isSuccess == "true" {
		fmt.Println("os.Exit(0)")
		os.Exit(0)
	}

	fmt.Println("os.Exit(1)")
	os.Exit(1)
}
