package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	fName string
	lName string
}

type secretAgent struct {
	person
	permissionToKill bool
}

func (p person) speak() {
	fmt.Println("Hello my name is ", p.fName, p.lName)
}

func (s secretAgent)  speak(){
	fmt.Println("Hello my name is ", s.lName, s.fName, s.lName)
}

func speak(h human) {
	h.speak()
}

func main() {
	p := person{"Leandro", "Rosa"}
	speak(p)

	s := secretAgent{person{"James", "Bond"}, true}
	speak(s)
}
