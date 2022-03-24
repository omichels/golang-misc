package main

import (
	"fmt"
	"golang-misc/fuentbuilder"
)

func main() {

	builder := fuentbuilder.NewEmployeeBuilder()

	peter := builder.Phone("555-12345678").Name("Peter Meier").Works("IBM").Salary("3500").Build()

	fmt.Println(peter)
}
