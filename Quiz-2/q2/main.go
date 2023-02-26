package main

import(
	"fmt"
)

func Age(value chan int){
	// Write to a channel
	value <- 21
	close(value)
	fmt.Println("Hello!!!")
}

func Name(name chan string){
	name <- "john doe"
	close(name)
}

func main(){
	// Create Channel
	value := make(chan int)
	name := make(chan string)
	// Goroutine
	go Age(value)
	go Name(name)
	// read from channel
	Age := <- value
	Name := <- name
	// Print the channel
	fmt.Println(Age, Name)
}