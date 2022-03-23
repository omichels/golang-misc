package main

import (
	"fmt"
)

type People interface {
	DoGreet() string
}

type GermanType struct {
	id int
}

type FrenchType struct {
	id int
}

func (g GermanType) DoGreet() string {
	return "Guten Tag"
}

func (f FrenchType) DoGreet() string {
	return "Bonjour"
}

func factory(language string, number int) People {
	factoryMap := map[string]func(int) People{
		"deutsch": func(i int) People {
			return &GermanType{
				id: i,
			}
		},
		"francais": func(i int) People {
			return &FrenchType{
				id: i,
			}
		},
	}
	newPerson := factoryMap[language](number)
	return newPerson
}

func main() {

	for i := 0; i < 4; i++ {
		var p People
		if i%2 == 0 {
			p = factory("deutsch", i)
		} else {
			p = factory("francais", i)
		}
		fmt.Print(i, " ")
		fmt.Println(p.DoGreet())
	}

}
