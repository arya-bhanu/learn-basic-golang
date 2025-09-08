package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	senderChannel := make(chan string)
	names := []string{
		"Amel",
		"Bahreisy",
		"Putu",
		"Shafira",
	}
	go func() {
		for _, name := range names {
			sendMessage(senderChannel, name)
		}
		// think senderChannel like consumer, you must tell him to close it's mouth if it's done/ running out of food stocks
		close(senderChannel)
	}()

	for by := range senderChannel {
		fmt.Println(by)
	}
}

func sendMessage(strChan chan<- string, from string) {
	strChan <- fmt.Sprintf("Hello from: %s\n", from)
}
