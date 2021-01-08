package main

import "fmt"

func main() {
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		defer fmt.Println("DONDONE")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Done")
}
