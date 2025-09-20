package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	generateFiles()
	duration := time.Since(start)
	fmt.Printf("generate files done in: %v\n", duration)
	start = time.Now()
	proceedSync()
	duration = time.Since(start)
	fmt.Printf("renaming files done in: %v\n", duration)
}
