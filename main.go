package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func Hello() string {
	return "Hello, world"
}

func Aggregate() string {
	return ""
}

func main() {
	// get file from terminal
	inputFile := os.Args[1]
	// declare chunk size
	maxSz, _ := strconv.Atoi(os.Args[2])
	// read the whole content of file and pass it to file variable, in case of error pass it to err variable
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	defer file.Close()

	// create buffer
	b := make([]byte, maxSz)

	for {
		// read content to buffer
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fileContent := string(b[:readTotal])
		// print content from buffer
		fmt.Println(fileContent)
	}
}

// SOURCE: https://www.golinuxcloud.com/golang-read-a-file-methods/
