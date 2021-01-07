package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Println("My favarite number is", rand.Intn(10))
	// 擬似乱数を生成するためにrand.Seed
	// 擬似乱数の生成にはよくtime.nowとunixNanoの組み合わせで実行できる
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favarite number is", rand.Intn(10))
}
