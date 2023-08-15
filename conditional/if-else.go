package main

import "fmt"

func main() {
	x := 10
	result := ""
	if x == 10 && x%2 == 0 {
		result = "less"
	} else {
		result = "more"
	}
	fmt.Println(result)

	// special syntex in golang where we can define variable in the if statemnet

	if y := 5; y == 5 && y%5 == 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
