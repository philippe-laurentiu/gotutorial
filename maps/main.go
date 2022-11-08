package main

import "fmt"

func main() {

	//doesent work
	var m1 map[string]string
	// m1["test"] = "valfortest"
	
	m2 := make(map[string]string)
	m2["test"] = "valufortest"
	delete(m2, "test")

	m3 := map[string]string{
		"red":  "#12333",
		"blue": "#8sfhsf",
	}
	m3["hase"] = "sdfsf"

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
}
