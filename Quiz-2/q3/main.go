package main

import (
	"fmt"
	"time"
)

func food(){
	fmt.Println("I like Apples")
}

func drink(){
	fmt.Println("I like Soda")
}
// A goroutine
func main(){
	go food()
	go drink()

	time.Sleep(time.Second)
}