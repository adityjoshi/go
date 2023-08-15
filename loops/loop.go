package main

import "fmt"

func main() {
	days := []string{"sunday", "monday", "wednesday", "thursday"}
	fmt.Println(days)
	// normal loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
		// advance loop
		for y := range days {
			fmt.Println(days[y])
		}
		// for each loop
	}
	// for each loop
	for index, day := range days {
		fmt.Printf("index is %v and the day is %v\n", index, day)
	}
	// while loop in go
	// how to use break and continue in go lang
	x := 1
	for x < 10 {
		if x == 5 {
			x++
			continue // it will just remove the number and add other nums
		}
		fmt.Println(x)
		x++
	}
	num := 1
	for num < 10 {
		if x == 8 {
			goto std
		}
		if num == 3 {
			break
		}
		fmt.Println(num)
		num++
	}
	// goto statement in golang in which when u hit the condition the statement will be printed
std:
	fmt.Println("the statement is hit when the condition is 8")
}
