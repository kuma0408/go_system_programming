package main

import (
	"io"
	"strings"
)

func main() {
	var reader io.Reader = strings.NewReader("test data")
	var readCloser io.ReadCloser = io.NopCloser(reader)
	readCloser.Close()
}
