package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type myWriter struct{}

func main() {
	res, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// two ways of writing it 
	// body := make([]byte, 99999)
	// res.Body.Read(body)
	// fmt.Println(string(body))

	// io.Copy(os.Stdout, res.Body)

	// with myWriter who implements Writer interface

	mw := myWriter{}
	io.Copy(mw, res.Body)
}

func (myWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("BYTES:",len(bs))
	return len(bs), nil
}
