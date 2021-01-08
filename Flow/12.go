package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("Hello")
	sayHello()
}

func sayHello() {

	defer fmt.Println("yaayyayaya")
	fmt.Println("HeyHelloooooo")
}
