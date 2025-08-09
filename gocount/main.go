package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Error: No filename provided.")
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: cannot read file")
		os.Exit(1)
	}

	fmt.Print(string(content))

}
