package main

import "fmt"

var zzz = 1000

// www := 2000

func main() {
	var i, j int = 1, 2

	// var xxx type = zzz の代わりに := を使うこともできる。
	// 関数の外では:=のような暗黙的宣言は使えず、varなどが必須
	k := 3
	c, python, java := true, false, "No!!"

	fmt.Println(zzz, i, j, k, c, python, java)
}
