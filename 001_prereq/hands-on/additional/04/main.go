package main

import "fmt"

type person struct {
	fName string
	lName string
}

func main() {
	p1 := person{"Leandro", "Rosa"}
	fmt.Println(p1.fName)
}
