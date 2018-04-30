package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname   string
	lastname    string
	contactInfo contactInfo
}

func main() {
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func log(msg string) {
	fmt.Println(msg)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

func (p person) updateName(newFirstName string) {
	p.firstname = newFirstName
}
