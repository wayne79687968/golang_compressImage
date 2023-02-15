package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter file name (support jpeg, png, gif): ")
	var filename string
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()
}
