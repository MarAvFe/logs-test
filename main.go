package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

const MAXMEMORYGB int = 1
const LOGSPATH string = "./logs/"

type LogLine struct {
	Date    time.Time
	Message string
}

func (l *LogLine) String() string {
	return l.Date.Format(time.RFC3339) + "," + l.Message
}

func (l *LogLine) After(other LogLine) bool {
	return l.Date.After(other.Date)
}

func (l *LogLine) Sort() {

}

func Hello() string {
	return "Hello, world"
}

func CloseOpenedFiles(files map[string]*os.File) {
	for _, fd := range files {
		fd.Close()
	}
}

func Aggregate() {

	// get values to avoid memory overflow. mem limits
	numFiles, chunkSize := getMaxChunkSize()

	// read all logs to aggregate
	files, err := ioutil.ReadDir(LOGSPATH)
	if err != nil {
		log.Fatal(err)
	}

	// create a structure to track files
	fileDescriptors := make(map[string]*os.File, numFiles)
	fileChunks := make(map[string]string, numFiles)
	logLines := make([]LogLine, 0)

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

	allDone := false
	for times := int64(0); !allDone; times++ {
		allDone = AllChunksNil(&fileChunks)
		loadChunks(fileDescriptors, fileChunks, chunkSize, times)
		for _, content := range fileChunks {
			//parsedLines := make([]LogLine, 0)
			for _, line := range strings.Split(content, "\n") {
				parts := strings.Split(line, ",")
				date, _ := time.Parse(time.RFC3339, parts[0])
				var message string
				if len(parts) > 1 {
					message = strings.Join(parts[1:], "")
				}
				logLines = append(logLines, LogLine{date, message})
			}
			//log.Println("parsed:", parsedLines)
			//loglines[fname] = parsedLines
			//log.Println(fname, content)
		}

		sort.Slice(logLines, func(i, j int) bool {
			return logLines[j].Date.After(logLines[i].Date)
		})

		for _, line := range logLines {
			fmt.Println(line.String())
		}

	}
}

func loadChunks(descriptors map[string]*os.File, files map[string]string, size int64, times int64) {
	//log.Println("loadChunks", descriptors, files, size, times)
	for fname, descriptor := range descriptors {
		buffer := make([]byte, size)

		// read content to buffer
		readTotal, err := descriptor.ReadAt(buffer, size*times)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
		}
		fileContent := string(buffer[:readTotal])
		if fileContent == "" { // File was already empty
			fileContent = "EOF"
		}
		// print content from buffer
		files[fname] = fileContent
		// SOURCE: https://www.golinuxcloud.com/golang-read-a-file-methods/
	}
}

func AllChunksNil(files *map[string]string) bool {
	//log.Println("\n\nallChunksNil", files)
	for _, chunk := range *files {
		if chunk[0:3] != "EOF" {
			return false
		}
	}
	return true
}

func getMaxChunkSize() (int, int64) {
	files, err := ioutil.ReadDir(LOGSPATH)
	if err != nil {
		log.Fatal(err)
	}

	numFiles := len(files)
	bytesinGB := 1000000000
	chunkSize := int64((MAXMEMORYGB * bytesinGB) / (numFiles / 2)) // div2 because of structured values
	return numFiles, chunkSize
}

func main() {

	//numFiles, chunkSize := getMaxChunkSize()
	//log.Printf("Max chunk size: %d bytes, which is %d max RAM GBs, divided into %d files \n", chunkSize, MAXMEMORYGB, numFiles)

	Aggregate()
}
