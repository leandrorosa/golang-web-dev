package main

import "fmt"

type person struct {
	fName string
	lName string
	favFood []string
}

func main() {
	p1 := person{"Leandro", "Rosa", []string{"steak", "pizza", "hambuger", "fries"}}
	fmt.Println(p1.favFood)
	for _,v := range p1.favFood {
		fmt.Println(v)
	}
}
