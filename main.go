package main

import (
	"fmt"
)

func main() {
	// var messages = make(chan string)

	// var sayHello = func(name string, iterate int) {
	// 	for i := range iterate {
	// 		fmt.Printf("%s:%d\n", name, i)
	// 	}
	// 	messages <- name
	// }

	// go sayHello("arnold", 10)
	// go sayHello("putu", 10)
	// go sayHello("andre", 10000)

	// fmt.Printf("Output first: %s\n", <-messages)
	// fmt.Printf("Output second: %s\n", <-messages)

	fmt.Println("=============================== CONTOH 2 =================================")

	var messagesChanel = make(chan string)
	names := []string{
		"izza",
		"arya",
		"angya",
	}
	for _, name := range names {
		go func(who string) {
			for i := range 100 {
				fmt.Printf("%s:%d\n", name, i)
			}
			messagesChanel <- who
		}(name)
	}

	var printMessage = func(messagesChanel chan string) {
		fmt.Printf("Hello from %s\n", <-messagesChanel)
	}

	for i := 0; i < len(names); i++ {
		printMessage(messagesChanel)
	}

}
