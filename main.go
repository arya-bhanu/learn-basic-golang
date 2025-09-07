package main

import (
	"fmt"
	"runtime"
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

	// fmt.Println("=============================== CONTOH 2 =================================")

	// var messagesChanel = make(chan string)
	// names := []string{
	// 	"izza",
	// 	"arya",
	// 	"angya",
	// }
	// for _, name := range names {
	// 	go func(who string) {
	// 		for i := range 100 {
	// 			fmt.Printf("%s:%d\n", name, i)
	// 		}
	// 		messagesChanel <- who
	// 	}(name)
	// }

	// var printMessage = func(messagesChanel chan string, isGoroutine string) {
	// 	fmt.Printf("Hello from %s and i'am from %s\n", <-messagesChanel, isGoroutine)
	// }

	// go printMessage(messagesChanel, "routine 1")

	// for i := 0; i < len(names)-1; i++ {
	// 	printMessage(messagesChanel, "")
	// }

	fmt.Println("==================================== CONTOH 3 =======================================")

	// buffered ni ibarat chanel memiliki kapasitas maksimal lebih dari 1. By default (un-buffered) memiliki kapasitas 1 sehingga akan bisa serah - terima satu-satu saja

	runtime.GOMAXPROCS(2)
	var messages = make(chan string, 3)

	persons := []string{
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
		"Amel",
		"Echi",
		"Hafid",
		"Hammuda",
		"Haqqy",
		"Arya",
		"Adek",
		"Restu",
	}

	var receiveMessage = func(messages chan string, index int) {
		fmt.Printf("Receive message: %s, %d\n", <-messages, index)
	}

	var sendMessage = func(message string, index int) {
		fmt.Printf("Sending message: %s, %d\n", message, index)
		messages <- message
	}

	// go routine saling terpisah-pisah tetap bisa mengakses chanel, karena akan saling asynchronus, harus ditunggu dengan fmt.Scanln
	// akan lancar berjalan karena asynchronus
	// for i := range len(persons) {
	// 	go receiveMessage(messages, i)
	// }
	// for i, person := range persons {
	// 	go sendMessage(person, i)
	// }
	// fmt.Scanln()

	// ini berbahaya ketika menerima duluan karena akan (sync), gaada go routine / pengiriman (gaakan bisa), program langsung selesai
	// for i := range len(persons) {
	// 	receiveMessage(messages, i)
	// }
	// for i, person := range persons {
	// 	go sendMessage(person, i)
	// }

	// ini akan bekerja, go routine akan berjalan terpisah untuk memasukan value pada chanel stream, tanpa perlu fmt.Scanln karena akan ditunggu oleh receiveMessage
	// ini tetap akan dikirim satu-satu, karena per loop akan ada go routine yang menerima - mengirim 1 buah value
	// for i, person := range persons {
	// 	go sendMessage(person, i)
	// }
	// for i := range len(persons) {
	// 	receiveMessage(messages, i)
	// }

	// ini works dengan buffers, tanpa harus menggunakan fmt.Scanln, karena penerima ada di main thread dan akan menunggu
	// go func() {
	// 	for i, person := range persons {
	// 		sendMessage(person, i)
	// 	}
	// }()

	// for i := range len(persons) {
	// 	receiveMessage(messages, i)
	// }

	// somehow channel apakah bisa benar-benar menunggu ya? meskipun terletak pada go routine yang berbeda? asal jumlahnya sesuai dengan berapa yang ingin kita terima dari pengirim
	// syaratnya, asalkan pengirim/penerima salah satunya ada di main-thread karena mereka dapat saling menunggu
	// go func() {
	// 	for i := range len(persons) {
	// 		receiveMessage(messages, i)
	// 	}
	// }()

	// for i, person := range persons {
	// 	sendMessage(person, i)
	// }

	// kalau seperti ini, harus menggunakan fmt.Scanln karena pengirim atau penerima gaada di main thread
	go func() {
		for i := range len(persons) {
			receiveMessage(messages, i)
		}
	}()
	go func() {
		for i, person := range persons {
			sendMessage(person, i)
		}
	}()
}
