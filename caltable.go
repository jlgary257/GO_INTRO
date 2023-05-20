package main

import "fmt"

func main() {

	var n int
	fmt.Printf("Please enter a number :")
	fmt.Scan(&n)
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(n, "X", i, "=", n*i)
		i++
	}
}
