package main

import(
	"fmt"
)

//Interface
type Vehicle interface{
	Car() string
}

// Use a struct to access the interface
type Type struct{
	make, model, brand string
}

// Use a struct to implement the Car() of the Interface
func (t Type) Car() string {
	return t.make + " " + t.model +   " " + t.brand
}

// Use a method to access the Interface
func vehicle(v Vehicle){
	fmt.Println("the Car is: ", v.Car())
}

func main(){
	// pass values to struct variables
	car1 := Type{
		"Mustang", 
		"Sedan",
		"ford",
	}

	// pass car1 to vehicle
	vehicle(car1)
}