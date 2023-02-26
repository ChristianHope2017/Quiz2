package main

import (
	"fmt"
	"sync"
)

func food(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("I like apples")
}

func drink(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("I like soda")
}

func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	go food(&wg)
	go drink(&wg)

	wg.Wait()
}