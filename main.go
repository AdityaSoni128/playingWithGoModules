package main

import (
	"fmt"

	calculator "github.com/AdityaSoni128/calculatorwithgo"
)

func main() {
	fmt.Println("Hello I am Playing with Go Modules")
	val := calculator.Add(10, 20)
	fmt.Println(val)
	calculator.Tablegetter(4)

}
