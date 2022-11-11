package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type writer struct{}

func main() {
	fn := os.Args[1]

	f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}

	w := writer{}

	io.Copy(w, f)

	if err := f.Close(); err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}
}

func (writer) Write(b []byte) (int, error) {
	fmt.Print(string(b))
	return len(b), nil
}
