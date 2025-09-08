package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("numbers :", numbers)

	ch1 := make(chan int)
	ch2 := make(chan float64)

	go getAverage(numbers, ch2)
	go getMax(numbers, ch1)

	for range 2 {
		select {
		case sum := <-ch1:
			fmt.Printf("Max value: %d\n", sum)
		case avg := <-ch2:
			fmt.Printf("Average: %f\n", avg)
		}
	}

}

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}
