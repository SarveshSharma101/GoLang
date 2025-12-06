package main

import (
	"fmt"
	"os"
)

func main() {
	file()
}

func file() {
	file := "./dummy.txt"
	fmt.Println("Stat a file")
	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Println("error stating a file: ", err)
	}
	if os.IsNotExist(err) {
		fmt.Println("Error file does now exist ", err)
	}
	fmt.Println("name: ", fileInfo.Name())
	fmt.Println("ModTime: ", fileInfo.ModTime())
	fmt.Println("Size: ", fileInfo.Size())
	fmt.Println("IsDir: ", fileInfo.IsDir())
	fmt.Println("Mode: ", fileInfo.Mode())

	fmt.Println("Opening a file....")

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("error opening the file: ", err)
	}
	fmt.Println("-->", f)
	defer f.Close()

	fmt.Println("reading a set bytes in file")
	b1 := make([]byte, 20)
	_, err = f.Read(b1)
	if err != nil {
		fmt.Println("error reading a file: ", err)
	}
	fmt.Println("-->", string(b1))

	fmt.Println("reading entire file")
	b, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error reading a file: ", err)
	}
	fmt.Println("-->", string(b))

	fmt.Println("Writing a file")
	nf, err := os.Create("./newFile.txt")
	if err != nil {
		fmt.Println("error opening the file: ", err)
	}
	fmt.Println("writing bytes to file")
	_, err = nf.Write([]byte("Hello world"))
	if err != nil {
		fmt.Println("error writing file:", err)
	}

	fmt.Println("writing string to file")
	_, err = nf.WriteString("Hello world again")
	if err != nil {
		fmt.Println("error writing file:", err)
	}
	defer nf.Close()

	fmt.Println("deleting the file")
	if err := os.Remove("./newFile.txt"); err != nil {
		fmt.Println("error deleting the file", err)
	}
}
