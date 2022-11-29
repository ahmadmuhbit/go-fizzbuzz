package main

import (
	"fmt"
	"testing"
)

func FizzBuzzUnitTest(total int) {
	for i := 1; i <= total; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func TestFizzBuzz2(t *testing.T) {
	FizzBuzzUnitTest(30)
}
