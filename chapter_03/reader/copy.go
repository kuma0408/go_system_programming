package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("file.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// io.Copy(os.Stdout, file)

	copy_file, err := os.Create("copy.go")
	if err != nil {
		panic(err)
	}
	defer copy_file.Close()
	io.Copy(copy_file, file)
}
