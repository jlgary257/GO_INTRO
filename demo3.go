package main

import "fmt"

func main() {
	a := 79
	b := 87
	fmt.Printf("a before : %d", a)
	//fmt.Println(a)
	fmt.Println()

	a = b
	fmt.Printf("a after :")
	fmt.Println(a)

	a += b
	fmt.Printf("a after addition : %d", a)

}
