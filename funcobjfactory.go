package main

import ( 
	"fmt"
)

type People interface{
		DoGreet() string
	}

type GermanType struct{
	id int
}

type FrenchType struc {
	id int
}

func  (g GermanType) DoGreet() string{
	return "Guten Tag"
}

func (f Frenchype) DoGreet() string{
	return "Bonjoir"
}
func InitGermanType(i int) People {
	return &GermanType{
	id: i
}
func InitFrenchType(i int) People {
	return &FrenchType{
			id: i
		}
func factory(language string, number int) People{

	factoryMap := map[string]func (int)  People{
		"german":InitGermanType,
		"french":InitFrenchType,
	}
	newPerson := factoryMap[language](i)

	return newPerson
}


func main() {

	for int i; i < 2; i++ {
			if i % 2 ==0 {
					fmt.Println(i)
					p:= factory("german",i)
					p.DoGreet()
				} else {
					fmt.Println(i)
					p:= factory("french",i)
					p.DoGreet()
				}

			}


