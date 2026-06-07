package main

import (
	"fmt"
	// "runtime"
)

func generateStream() <-chan int {
	str := make(chan int)
	go func() {
		for i := 1; i <= 1000; i++ {
			str <- i
		}
		close(str)
	}()

	return str

}

func main() {
	str := generateStream()
	prime := primeNumCheck(str)

	for val := range prime {
		fmt.Println(val)
	}

}

func primeNumCheck(str <-chan int) <-chan int {
	primeChan := make(chan int)
	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}
		half := num / 2
		for i := 2; i <= half; i++ {
			mod := num % i
			if mod == 0 {
				return false
			}
		}
		return true
	}

	go func() {
		for val := range str {
			if isPrime(val) {
				primeChan <- val
			}
		}
		close(primeChan)
	}()

	return primeChan

}
