package main

import (
	"fmt"

	"github.com/arya-bhanu/arya-go-package/webdev"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(quote.Go())
	res, err := webdev.RetryFuncHandler(3, func(a ...any) (any, error) {
		return "Hello world", nil
	})
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Printf("Result %+v\n", res)
}
