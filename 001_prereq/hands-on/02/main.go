package main

import "fmt"

type person struct {
	fName string
	lName string
}

type secretAgent struct {
	person
	permissionToKill bool
}

func (p person) pSpeak() {
	fmt.Println("Hello my name is ", p.fName, p.lName)
}

func (s secretAgent)  saSpeak(){
	fmt.Println("Hello my name is ", s.lName, s.fName, s.lName)
}

func main() {
	p := person{"Leandro", "Rosa"}
	p.pSpeak()

	s := secretAgent{person{"James", "Bond"}, true}
	s.saSpeak()
	s.pSpeak()
}
