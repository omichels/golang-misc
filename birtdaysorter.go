package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

func main() {

	people := []Person{
		Person{
			name: "Leela",
			age:  30,
		},
		Person{
			name: "Fry",
			age:  33,
		},
		Person{
			name: "Bender",
			age:  4,
		},
		Person{
			name: "Prof. Farnsworth",
			age:  160,
		}}

	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println(people)

}
