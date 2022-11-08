package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	test := person{
		firstName: "alex",
		lastName: "hase",
		contact: contactInfo{
			email: "a@a.de",
			zip: 123,
		},
	}
	test.updateName("moppel")
	test.print()
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}


func (p person) print() {
	fmt.Printf("%+v",p)
}