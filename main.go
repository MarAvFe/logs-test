package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const MAXMEMORYGB int = 1
const LOGSPATH string = "./logs/"

func Hello() string {
	return "Hello, world"
}

func CloseOpenedFiles(files map[string]*os.File) {
	for fname, fd := range files {
		log.Println("Closing file", fname)
		fd.Close()
	}
}

func Aggregate() string {

	// get values to avoid memory overflow. mem limits
	numFiles, _ := getMaxChunkSize()

	// read all logs to aggregate
	files, err := ioutil.ReadDir(LOGSPATH)
	if err != nil {
		log.Fatal(err)
	}

	// create a structure to track files
	fileDescriptors := make(map[string]*os.File, numFiles)

	// open all files and track its file descriptors
	for _, fname := range files {
		file, err := os.Open(LOGSPATH + fname.Name())
		if err != nil {
			fmt.Printf("Could not open the file due to this %s error \n", err)
			continue
		}
		fileDescriptors[fname.Name()] = file
	}
	defer CloseOpenedFiles(fileDescriptors) // Protect against memory leaks

	return ""
}

func getMaxChunkSize() (int, int64) {
	files, err := ioutil.ReadDir(LOGSPATH)
	if err != nil {
		log.Fatal(err)
	}

	numFiles := len(files)
	bytesinGB := 1000000000
	chunkSize := int64((MAXMEMORYGB * bytesinGB) / numFiles)
	return numFiles, chunkSize
}

func main() {

	numFiles, chunkSize := getMaxChunkSize()
	log.Printf("Max chunk size: %d bytes, which is %d max RAM GBs, divided into %d files \n", chunkSize, MAXMEMORYGB, numFiles)

	// get file from terminal
	inputFile := "./logs/server-bc329xbv.log"
	// declare chunk size
	maxSz, _ := strconv.Atoi("10")
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

	Aggregate()
}

// SOURCE: https://www.golinuxcloud.com/golang-read-a-file-methods/
