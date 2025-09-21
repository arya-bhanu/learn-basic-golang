package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	GenerateFileAsync()
	duration := time.Since(start)
	fmt.Printf("generate files done in: %v\n", duration)
	start = time.Now()
	// ProceedAsync()
	// proceedSync()
	duration = time.Since(start)
	fmt.Printf("renaming files done in: %v\n", duration)
}
