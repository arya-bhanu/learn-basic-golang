package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(3)
	go print(10, "halo")
	go print(10, "apa")
	print(10, "apa kabar")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)
}

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println(i+1, message)
	}
}
