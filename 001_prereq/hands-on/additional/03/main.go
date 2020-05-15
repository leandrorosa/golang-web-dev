package main

import "fmt"

type person struct {
	fName string
	lName string
}

func main() {
	p := person{"Leandro", "Rosa"}
	fmt.Println(p)
}
