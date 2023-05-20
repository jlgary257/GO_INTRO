package main

import "fmt"

func Fib(num int) {
	a := 0
	b := 1
	c := b
	fmt.Printf("Fib series is %d%d", a, b)
	for true {
		c = b
		b = a + b
		if b > num {
			fmt.Println()
			break
		}
		a = c
		fmt.Println(b)
	}
}

func main() {
	Fib(16)

}
