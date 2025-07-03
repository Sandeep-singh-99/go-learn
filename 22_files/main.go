package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	fileInfo, err := f.Stat()

	if err != nil {
		panic(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("File size:", fileInfo.Size())
	fmt.Println("File mode:", fileInfo.Mode())
	fmt.Println("File modification time:", fileInfo.ModTime())
	fmt.Println("Is directory:", fileInfo.IsDir())

	defer f.Close()
	fmt.Println("File closed successfully")
}