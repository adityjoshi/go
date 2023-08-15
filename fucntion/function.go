package main

import "fmt"

func main() {
	name("adi")
	// print array
	var ages = [2]string{"hey", "why"}
	names := [3]int{11, 12, 21}
	fmt.Print(names, ages)
	// print slices
	var fruits = []string{"apple", "mango", "papaya", "banana"}
	fmt.Println(fruits)
	//to add something in slices
	fruits = append(fruits, "peach", "orange")
	fmt.Println(fruits)
	fruits = append(fruits[1:]) // slice will print from x [x:] x is inclusive
	fmt.Println(fruits)
	fruits = append(fruits[1:3]) // [1:3] 1 will be inclusive and 3 will not be incusive so it will go from 1 to 2 only
	// one of the best thing about golang is that it realcotes the memory while you append something in array it
	// helps in saving a lot of memory.

	highscore := make([]int, 4)
	highscore[0] = 1
	highscore[1] = 2
	highscore[2] = 3
	highscore[3] = 4
	highscore = append(highscore, 233, 444)
	fmt.Println(highscore)
	// how to remove a element from the list using index
	var courses = []string{"reactjs", "swift", "java", "golang", "ruby"}
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

func name(s string) {
	fmt.Print("hey \n", s)
}
