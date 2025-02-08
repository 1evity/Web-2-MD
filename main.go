package main

import (
	"fmt"
    "os"
    "log"
    "./internal/declutter"
    "./internal/convert"
	"./internal/save"
)

func main() {
	// Check command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: w2md <link-to-webpage>")
		os.Exit(1)
	}

	// Input webpage URL
	pageURL := os.Args[1]
	
}