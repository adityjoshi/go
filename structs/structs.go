package main

import "fmt"

func main() {
	// there is no inheritance in golang no super or parent class
	adi := User{"aditya", 18, false, "adi@30joshi"}
	fmt.Println(adi)
	// to print info in a more efficent way
	fmt.Printf("adi's details are: %+v\n", adi)
	// for some specific info
	fmt.Printf("name is %v and age is %v \n", adi.Name, adi.Age)
}

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}
