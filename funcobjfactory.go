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
func InitGermanType(i int) People {
	return &GermanType{
		id: i,
	}
}
func InitFrenchType(i int) People {
	return &FrenchType{
		id: i,
	}
}
func factory(language string, number int) People {

	factoryMap := map[string]func(int) People{
		"german": InitGermanType,
		"french": InitFrenchType,
	}
	newPerson := factoryMap[language](number)
	return newPerson
}

func main() {

	for i := 0; i < 2; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
			p := factory("german", i)
			fmt.Println(p.DoGreet())
		} else {
			fmt.Print(i, " ")
			p := factory("french", i)
			fmt.Println(p.DoGreet())
		}

	}

}
