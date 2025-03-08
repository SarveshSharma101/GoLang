package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// create a file
	filePath := "temp.txt"
	var file *os.File
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err = os.Create(filePath)
		if err != nil {
			log.Fatalf("error %v creating the file %v", err, filePath)
		}
	} else if err != nil {
		log.Fatalf("error %v while trying to stat the file %v", err, filePath)
	} else {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Fatalf("error %v opening the file %v", err, filePath)
		}
	}
	defer file.Close()
	_, err := file.WriteString("Hellow world")
	if err != nil {
		log.Fatalf("error %v writing to file %v", err, filePath)
	}

	r, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error %v reading the file %v", err, filePath)
	}
	fmt.Println("read ", string(r))
}
