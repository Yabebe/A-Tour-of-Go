package main

import "fmt"

func main() {
	sum := 0
	// for 初期ステートメント ; 条件式 ; 後処理ステートメント { }
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		fmt.Println("Added 1 to i")
		sum += i
		for i := 0; i < 10 && sum < 45; i++ {
			fmt.Println("Still looping")
		}
	}
	fmt.Println(sum)
}
