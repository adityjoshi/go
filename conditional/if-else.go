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
	fmt.Print(result)
}
