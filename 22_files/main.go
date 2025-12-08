package main

import (
	// "bufio"
	"fmt"
	"os"
)

// func main() {
// 	f, err := os.Open("22_files/example.txt")

// 	if err != nil {
// 		panic(err)
// 	}

// 	fileInfo, err := f.Stat()

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("File name:", fileInfo.Name())
// 	fmt.Println("File size:", fileInfo.Size())
// 	fmt.Println("File mode:", fileInfo.Mode())
// 	fmt.Println("File modification time:", fileInfo.ModTime())
// 	fmt.Println("Is directory:", fileInfo.IsDir())

// 	defer f.Close()
// 	fmt.Println("File closed successfully")
// }


func main() {
	// f, err := os.Open("22_files/example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 10)
	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }

	// for i:=0; i<len(buf); i++ {
	// 	fmt.Println("data", d, string(buf[i]))
	// }

	// fmt.Println("data", d, buf)

	// read file using ReadFile
	// f, err := os.ReadFile("22_files/example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("data", string(f))

	// read folders

	// dir, err := os.Open(".")

	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.Readdir(-1)

	// if err != nil {
	// 	panic(err)
	// }

	// for _ , fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir(), fi.Size(), fi.Mode())
	// }

	// Create a file
	// f, err := os.Create("22_files/newfile.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("Hi, I am Sandeep Singh")

	// fmt.Println("File created successfully")

	// read and write to another file (streaming fashion)

	// srcFile, err := os.Open("22_files/example.txt")
	
	// if err != nil {
	// 	panic(err)
	// }

	// defer srcFile.Close()

	// destFile, err := os.Create("22_files/destfile.txt")
	
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(srcFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()

	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("File copied successfully")

	// delete a file


	err := os.Remove("22_files/newfile.txt")
	
	if err != nil {
		panic(err)
	}

	fmt.Println("File deleted successfully")
}