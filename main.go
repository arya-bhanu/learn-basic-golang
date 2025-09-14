package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	greetingsChan := make(chan string, 1)
	letterLengthChan := make(chan int)
	friends := []string{
		"Bahreisy",
		"Amel",
		"Shafira",
		"Restu",
	}

	go func() {
		for _, friend := range friends {
			fmt.Printf("Send message: %s\n", friend)
			sendMessage(greetingsChan, friend)
		}
		close(greetingsChan)
	}()

	go func() {
		for _, friend := range friends {
			countLetterName(letterLengthChan, friend)
		}
		close(letterLengthChan)
	}()

	// for range friends {
	// 	select {
	// 	case greet := <-greetingsChan:
	// 		fmt.Println(greet)
	// 	case <-time.After(time.Second * 5):
	// 		fmt.Println("Time out sendMessage")
	// 		continue
	// 	}
	// }

	go func() {
		for greet := range greetingsChan {
			fmt.Println(greet)
		}
	}()

	sisa, ok := <-greetingsChan
	if ok {
		fmt.Printf("SISA NAMA: %s\n", sisa)
	}

	for range friends {
		select {
		case countLetters := <-letterLengthChan:
			fmt.Printf("Count letters %d\n", countLetters)
		case <-time.After(time.Second * 5):
			fmt.Println("Time out countLetterName")
			continue
		}
	}

	fmt.Scanln()
}

func sendMessage(strChan chan<- string, from string) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	// only for sending only into chanel
	strChan <- from
}

func countLetterName(letterLengthChan chan<- int, from string) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)

	letterLengthChan <- len(from)
}
