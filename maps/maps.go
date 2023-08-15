package main

import (
	"fmt"
)

func main() {
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["jdtls"] = "java"
	languages["vi"] = "neovim"
	fmt.Println(languages)
	fmt.Println("js : ", languages["js"])

	// for loops in maps
	for key, value := range languages { // if i put _ in place of key it will just remove the key
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

	// to delete something from the maps
	delete(languages, "jdtls")
	fmt.Println(languages)
}
